package handler

import (
	"encoding/json"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/errno"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/jwt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/response"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/query"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/autoPlan"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/caseAssemble"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/homePage"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/machine"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/plan"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/report"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/scene"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/target"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/uiReport"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/uiScene"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/variable"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/public"
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"net/http"
	"strings"
	"sync"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024000,
	WriteBufferSize: 1024000,
}

// WebSocketReq wb消息结构体
type WebSocketReq struct {
	RouteUrl string `json:"route_url" binding:"required"`
	Param    string `json:"param" binding:"required"`
}

// ClientLinkList 客户端连接信息
var ClientLinkList = make(map[string]*ClientLink)

// UIEngineTopicMap ui 自动化topic
var UIEngineTopicMap = make(map[string]struct{})

// ClientLink 链接基本数据结构
type ClientLink struct {
	TeamID    string          `json:"team_id"`
	LinkTime  int64           `json:"link_time"`
	Token     string          `json:"token"`
	Websocket *websocket.Conn `json:"websocket"`
	Mu        sync.Mutex
}

// ClientTeamAndUserMap 客户端连接团队与用户的映射关系
var ClientTeamAndUserMap = make(map[string][]string)

// CloseInvalidWbLink CloneLink 关闭无用的链接
func CloseInvalidWbLink() {
	for {
		nowTime := time.Now().Unix()
		for userID, linkInfo := range ClientLinkList {
			if linkInfo.LinkTime < nowTime-80 {
				err := linkInfo.Websocket.Close()
				if err != nil {
					log.Logger.Error("关闭wb链接--失败，err:", err)
				}
				delete(ClientLinkList, userID) // 删除客户端连接信息
			}
		}
		time.Sleep(80 * time.Second)
	}
}

// ConsumerUIEngineResult 消费 UI 自动化结果
func ConsumerUIEngineResult() {
	for {
		ctx := context.Background()
		keys, err := dal.GetRDB().Keys(ctx, "UiReport:*").Result()
		if err != nil {
			log.Logger.Error("ConsumerUIEngineResult--dal.GetRDB().Scan，err:", proof.WithError(err))
			time.Sleep(5 * time.Second)
			continue
		}
		for _, key := range keys {
			// 一个 go 消费一个队列
			if _, ok := UIEngineTopicMap[key]; !ok {
				UIEngineTopicMap[key] = struct{}{}
				go func(k string) {
					handleUIEngineResult(ctx, k)
				}(key)
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func handleUIEngineResult(ctx context.Context, key string) {
	for {
		// 设置一个5秒的超时时间
		dataList, err := dal.GetRDB().BRPop(ctx, 5*time.Second, key).Result()
		if err == redis.Nil {
			continue
		}
		if err != nil {
			log.Logger.Error("handleUIEngineResult BRPop err:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		log.Logger.Info("ConsumerUIEngineResult，dataList:", dataList)
		// 查询到了数据
		var resultData *rao.UIEngineResultDataMsg
		if len(dataList) == 2 {
			resultMsgString := dataList[1]

			err = json.Unmarshal([]byte(resultMsgString), &resultData)
			if err != nil {
				log.Logger.Info("ConsumerUIEngineResult--json转换格式错误，err:", proof.WithError(err))
				continue
			}

			// 查询是否存在报告
			pr := query.Use(dal.DB()).UIPlanReport
			planReport, _ := pr.WithContext(ctx).Where(pr.ReportID.Eq(resultData.TopicID)).First()
			if planReport != nil {
				resultData.IsReport = true
			}

			// 判断任务是否已经停止
			redisKey := consts.UIEngineRunAddrPrefix + resultData.TopicID
			addr, err := dal.GetRDB().Get(ctx, redisKey).Result()
			if err != nil {
				log.Logger.Error("ConsumerUIEngineResult--Get UIEngineRunAddrPrefix err:", err)
			}
			if len(addr) == 0 {
				continue
			}

			var withdraws = make([]*rao.UIEngineDataWithdraw, 0, len(resultData.DataWithdraws))
			for _, withdraw := range resultData.DataWithdraws {
				withdraws = append(withdraws, withdraw)
			}
			resultData.Withdraws = withdraws

			collectionName := consts.CollectUISendSceneOperator
			if !resultData.IsReport {
				collectionName = consts.CollectUISendSceneOperatorDebug
			}
			collection := dal.GetMongo().Database(dal.MongoDB()).Collection(collectionName)
			var teamID string
			// End 单独结构
			if resultData.End {
				// 删除 Redis 对应关系
				dal.GetRDB().SRem(ctx, consts.UIEngineCurrentRunPrefix+addr, resultData.TopicID)
				dal.GetRDB().Del(ctx, consts.UIEngineRunAddrPrefix+resultData.TopicID)

				filter := bson.D{{"report_id", resultData.TopicID}}
				uiSendSceneOperator := &mao.UISendSceneOperator{}
				if err = collection.FindOne(ctx, filter).Decode(&uiSendSceneOperator); err != nil {
					log.Logger.Error("ConsumerUIEngineResult--FindOne err:", err)
				}
				resultData.SceneID = uiSendSceneOperator.SceneID
				teamID = uiSendSceneOperator.TeamID

				// 处理报告结束
				if err = uiReport.HandleReportEnd(ctx, resultData); err != nil {
					log.Logger.Error("HandleReportEnd--处理报告处理，SceneID:", resultData.SceneID, "report ID:", resultData.TopicID, "err:", err)
				}
			} else {
				// 组织断言结果数据
				filter := bson.D{
					{"report_id", resultData.TopicID},
					{"scene_id", resultData.SceneID},
					{"operator_id", resultData.OperatorID}}
				uiSendSceneOperator := &mao.UISendSceneOperator{}
				if err = collection.FindOne(ctx, filter).Decode(&uiSendSceneOperator); err != nil {
					log.Logger.Error("ConsumerUIEngineResult--FindOne err:", err)
				}
				var a *mao.AssertResults
				if err := bson.Unmarshal(uiSendSceneOperator.AssertResults, &a); err != nil {
					log.Logger.Errorf("uiSendSceneOperator.AssertResults bson unmarshal err %w", err)
				}
				if len(a.Asserts) == len(resultData.Assertions) {
					for k, assertion := range resultData.Assertions {
						assertion.Name = a.Asserts[k].Name
					}
				}

				// 是否有多组数据
				if len(uiSendSceneOperator.ParentID) > 1 {
					so := query.Use(dal.DB()).UISceneOperator
					parentScene, _ := so.WithContext(ctx).Where(so.OperatorID.Eq(uiSendSceneOperator.ParentID)).First()
					if parentScene != nil {
						if parentScene.Action == consts.UISceneOptTypeForLoop || parentScene.Action == consts.UISceneOptTypeWhileLoop {
							resultData.IsMulti = true
						}
						// for 步骤默认成功
						if parentScene.Action == consts.UISceneOptTypeForLoop {
							_ = uiScene.UpdateSimpleParent(ctx, resultData, parentScene.OperatorID)
						}
					}
				}

				uiEngineResultDataMsg, err := uiScene.UpdateReportDataResultMulti(ctx, resultData)
				if err != nil {
					log.Logger.Error("UpdateReportData--修改报告数据失败，SceneID:", resultData.SceneID, "userID:", resultData.UserID, "err:", err)
				}
				resultData.MultiResult = uiEngineResultDataMsg
				teamID = uiSendSceneOperator.TeamID
			}

			// 查询用户连接
			if l, ok := ClientLinkList[resultData.UserID]; ok {
				l.Mu.Lock()
				defer l.Mu.Unlock()

				resp := response.WbSuccessWithData(ctx, resultData, "ui_engine_result")
				err = l.Websocket.WriteMessage(websocket.TextMessage, []byte(resp))
				if err != nil {
					log.Logger.Error("ConsumerUIEngineResult--给用户发送消息失败，SceneID:", resultData.SceneID, "userID:", resultData.UserID, "err:", err)
					continue
				}
			}

			// 报告推团队 场景推用户
			if resultData.IsReport {
				if userIDs, ok := ClientTeamAndUserMap[teamID]; ok {
					for _, userID := range userIDs {
						if l, ok := ClientLinkList[userID]; ok && userID != resultData.UserID {
							l.Mu.Lock()
							defer l.Mu.Unlock()

							resp := response.WbSuccessWithData(ctx, resultData, "ui_engine_result")
							err = l.Websocket.WriteMessage(websocket.TextMessage, []byte(resp))
							if err != nil {
								log.Logger.Error("ConsumerUIEngineResult--给用户发送消息失败，SceneID:", resultData.SceneID, "userID:", resultData.UserID, "err:", err)
							}
						}
					}
				}
			}

			if resultData.End {
				goto end // 跳出两层循环
			}

		}
		time.Sleep(5 * time.Second)
	}
end: // 定义一个标签
	log.Logger.Info("ConsumerUIEngineResult end:", key)
	delete(UIEngineTopicMap, key)
}

// PushRunningPlanCount 主动推送运行中计划数量的方法
func PushRunningPlanCount() {
	msgType := 1
	for {
		if len(ClientLinkList) > 0 {
			// 组装团队与团队所属用户的信息
			teamAndUserMap := make(map[string][]string, len(ClientLinkList))
			for userID, linkInfo := range ClientLinkList {
				teamAndUserMap[linkInfo.TeamID] = append(teamAndUserMap[linkInfo.TeamID], userID)
			}

			// 查询正在运行的性能计划
			for teamID, userArr := range teamAndUserMap {
				ctx := context.Background()
				runningPlanNum, err := plan.ListByStatus(ctx, teamID)
				if err != nil {
					log.Logger.Info("运行中计划--查询失败，err:", err)
					break
				}
				respData := rao.ListUnderwayPlanResp{
					RunPlanNum: runningPlanNum,
				}
				resp := response.WbSuccessWithData(ctx, respData, "running_plan_count")

				// 给这个团队下所有用户发送消息
				for _, userID := range userArr {
					// 写入ws数据
					if websocketLink, ok := ClientLinkList[userID]; ok {
						websocketLink.Mu.Lock()
						defer websocketLink.Mu.Unlock()

						err = websocketLink.Websocket.WriteMessage(msgType, []byte(resp))
						if err != nil {
							log.Logger.Info("运行中计划--给用户发送消息失败，teamID:", teamID, "userID:", userID, "err:", err)
							break
						}
					}
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// WebSocket 建立websocket连接方法
func WebSocket(ctx *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Logger.Info("websocket建立链接失败：", err)
		return
	}
	log.Logger.Info("建立链接成功")

	for {
		// 读取ws中的数据
		msgType, message, err := ws.ReadMessage()
		log.Logger.Info("接受消息--原始消息：", msgType, string(message))
		if err != nil {
			log.Logger.Info("接受消息失败：", err)
			break
		}

		var webSocketReq WebSocketReq
		err = json.Unmarshal(message, &webSocketReq)
		if err != nil {
			break
		}

		log.Logger.Info("接受消息--结构化：", webSocketReq)

		// 分发逻辑
		res, err := Dispense(ctx, ws, &webSocketReq)
		if res != "" {
			err = ws.WriteMessage(msgType, []byte(res))
			if err != nil {
				break
			}
		}
	}
}

// Dispense 分发请求
func Dispense(ctx *gin.Context, ws *websocket.Conn, wbReq *WebSocketReq) (string, error) {
	respString := ""
	var err error

	// 选择路由
	switch wbReq.RouteUrl {
	case "start_heartbeat":
		respString, err = StartHeartbeat(ctx, wbReq, ws)
	case "save_scene_flow":
		respString, err = SaveSceneFlow(ctx, wbReq)
	case "save_case_flow":
		respString, err = SaveCaseFlow(ctx, wbReq)
	case "home_page":
		respString, err = HomePageData(ctx, wbReq)
	case "stress_plan_list":
		respString, err = StressPlanList(ctx, wbReq)
	case "stress_report_list":
		respString, err = StressReportList(ctx, wbReq)
	case "auto_plan_list":
		respString, err = AutoPlanList(ctx, wbReq)
	case "auto_plan_detail":
		respString, err = AutoPlanDetail(ctx, wbReq)
	case "auto_report_list":
		respString, err = AutoReportList(ctx, wbReq)
	case "stress_plan_detail":
		respString, err = StressPlanDetail(ctx, wbReq)
	//case "stress_report_debug":
	//	respString, err = StressReportDebug(ctx, wbReq)
	case "stress_report_machine_monitor":
		respString, err = StressReportMachineMonitor(ctx, wbReq)
	case "stress_report_task_detail":
		respString, err = StressReportTaskDetail(ctx, wbReq)
	case "stress_report_detail":
		respString, err = StressReportDetail(ctx, wbReq)
	case "stress_machine_list":
		respString, err = StressMachineList(ctx, wbReq)
	case "send_api_result":
		respString, err = SendApiResult(ctx, wbReq)
	case "send_scene_result":
		respString, err = SendSceneResult(ctx, wbReq)
	case "user_switch_team":
		UserSwitchTeam(ctx, wbReq, ws)
	case "disband_team_notice":
		DisbandTeamNotice(ctx, wbReq)
	case "save_global_param":
		respString, err = SaveGlobalParam(ctx, wbReq)
	case "get_global_param":
		respString, err = GetGlobalParam(ctx, wbReq)
	case "save_scene_param":
		respString, err = SaveSceneParam(ctx, wbReq)
	case "get_scene_param":
		respString, err = GetSceneParam(ctx, wbReq)
	default:
		respData := response.Response{}
		resTemp, _ := json.Marshal(respData)
		respString, err = string(resTemp), nil
	}
	return respString, err
}

type StartHeartbeatReq struct {
	Token  string `json:"token"`
	TeamID string `json:"team_id"`
}

type UserSwitchTeamReq struct {
	Token  string `json:"token"`
	TeamID string `json:"team_id"`
}

func StartHeartbeat(ctx *gin.Context, wbReq *WebSocketReq, ws *websocket.Conn) (string, error) {
	var startHeartbeatReq StartHeartbeatReq
	err := json.Unmarshal([]byte(wbReq.Param), &startHeartbeatReq)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	userID, err := jwt.ParseToken(startHeartbeatReq.Token)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrInvalidToken, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	// 判断当前用户id是否建联过
	if clientLink, ok := ClientLinkList[userID]; ok {
		// 判断token是否相同
		if startHeartbeatReq.Token != clientLink.Token {
			// 给当前用户之前的连接发送断连退出登录消息
			msgType := 1
			resp := response.WbSuccess(ctx, "user_logout")
			err = clientLink.Websocket.WriteMessage(msgType, []byte(resp))
		}
	}

	// 保存用户心跳数据,把客户端链接加入到链接池当中
	nowTime := time.Now().Unix()
	ClientLinkList[userID] = &ClientLink{
		TeamID:    startHeartbeatReq.TeamID,
		LinkTime:  nowTime,
		Token:     startHeartbeatReq.Token,
		Websocket: ws,
	}

	// 保存团队与团队里面所属用户的映射关系
	needAdd := 1
	for _, teamUserID := range ClientTeamAndUserMap[startHeartbeatReq.TeamID] {
		if userID == teamUserID {
			needAdd = 0
		}
	}
	// 把用户id添加到所属团队下面
	if needAdd == 1 {
		ClientTeamAndUserMap[startHeartbeatReq.TeamID] = append(ClientTeamAndUserMap[startHeartbeatReq.TeamID], userID)
	}

	resp := response.WbSuccess(ctx, wbReq.RouteUrl)
	return resp, nil
}

// UserSwitchTeam 用户切换团队
func UserSwitchTeam(ctx *gin.Context, wbReq *WebSocketReq, ws *websocket.Conn) {
	var userSwitchTeamReq UserSwitchTeamReq
	err := json.Unmarshal([]byte(wbReq.Param), &userSwitchTeamReq)
	if err != nil {
		log.Logger.Info("用户切换团队--参数错误")
	}
	userID, err := jwt.ParseToken(userSwitchTeamReq.Token)
	if err != nil {
		log.Logger.Info("用户切换团队--用户token解析失败")
	}

	// 保存用户心跳数据,把客户端链接加入到链接池当中
	nowTime := time.Now().Unix()
	ClientLinkList[userID] = &ClientLink{
		TeamID:    userSwitchTeamReq.TeamID,
		LinkTime:  nowTime,
		Token:     userSwitchTeamReq.Token,
		Websocket: ws,
	}
}

func SaveSceneFlow(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.SaveFlowReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	log.Logger.Info("保存场景flow参数结构体：", req)

	for _, nodeInfo := range req.Nodes {
		if nodeInfo.Type == consts.FlowTypeWaitController && nodeInfo.WaitMs > 20000 {
			resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
			return resp, nil
		}
	}

	errNum, err := scene.SaveFlow(ctx, &req)
	log.Logger.Info("保存场景flow返回值：", errNum, err)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errNum, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccess(ctx, wbReq.RouteUrl)
	return resp, nil
}

func SaveCaseFlow(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.SaveSceneCaseFlowReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	log.Logger.Info("保存用例flow参数结构体：", req)

	for _, nodeInfo := range req.Nodes {
		if nodeInfo.Type == consts.FlowTypeWaitController && nodeInfo.WaitMs > 20000 {
			resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
			return resp, nil
		}
	}

	err = caseAssemble.SaveSceneCaseFlow(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccess(ctx, wbReq.RouteUrl)
	return resp, nil
}

func HomePageData(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.HomePageReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	homePageData, err := homePage.HomePage(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccessWithData(ctx, homePageData, wbReq.RouteUrl)
	return resp, nil
}

func StressPlanList(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.ListPlansReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	plans, total, err := plan.ListByTeamID(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.ListPlansResp{
		Plans: plans,
		Total: total,
	}

	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func StressReportList(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.ListReportsReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	isExist := strings.Index(req.Keyword, "%")
	if isExist >= 0 {
		resp := response.WbSuccessWithData(ctx, rao.ListReportsResp{}, wbReq.RouteUrl)
		return resp, err
	}

	reports, total, err := report.GetReportList(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.ListReportsResp{
		Reports: reports,
		Total:   total,
	}

	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func AutoPlanList(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetAutoPlanListReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	isExist := strings.Index(req.PlanName, "%")
	if isExist >= 0 {
		resp := response.WbSuccessWithData(ctx, rao.AutoPlanListResp{}, wbReq.RouteUrl)
		return resp, err
	}

	list, total, err := autoPlan.GetAutoPlanList(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.AutoPlanListResp{
		AutoPlanList: list,
		Total:        total,
	}
	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func AutoPlanDetail(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetAutoPlanDetailReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	detail, err := autoPlan.GetAutoPlanDetail(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccessWithData(ctx, detail, wbReq.RouteUrl)
	return resp, nil
}

func AutoReportList(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetAutoPlanReportListReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	isExist := strings.Index(req.PlanName, "%")
	if isExist >= 0 {
		resp := response.WbSuccessWithData(ctx, rao.GetAutoPlanReportListResp{}, wbReq.RouteUrl)
		return resp, err
	}

	list, total, err := autoPlan.GetAutoPlanReportList(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.GetAutoPlanReportListResp{
		AutoPlanReportList: list,
		Total:              total,
	}
	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func StressPlanDetail(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetPlanConfReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	p, err := plan.GetByPlanID(ctx, req.TeamID, req.PlanID)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.GetPlanResp{Plan: p}
	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

//func StressReportDebug(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
//	var req rao.GetReportReq
//	err := json.Unmarshal([]byte(wbReq.Param), &req)
//	if err != nil {
//		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
//		return resp, err
//	}
//
//	result, err := report.GetReportDebugLog(ctx, req)
//	if err != nil {
//		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
//		return resp, err
//	}
//
//	resp := response.WbSuccessWithData(ctx, result, wbReq.RouteUrl)
//	return resp, nil
//}

func StressReportMachineMonitor(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.ListMachineReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	machineData, err := report.ListMachines(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	resp := response.WbSuccessWithData(ctx, machineData, wbReq.RouteUrl)
	return resp, nil
}

func StressReportTaskDetail(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetReportTaskDetailReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	ret, err := report.GetTaskDetail(ctx, req)
	if err != nil {
		if err.Error() == "报告不存在" {
			resp := response.WbErrorWithMsg(ctx, errno.ErrReportNotFound, err.Error(), wbReq.RouteUrl)
			return resp, err
		} else {
			resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
			return resp, err
		}
	}

	respData := rao.GetReportTaskDetailResp{Report: ret}

	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func StressReportDetail(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetReportReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	result, err := report.GetReportDetail(ctx, req)
	if err != nil {
		if err.Error() == "报告不存在" {
			resp := response.WbErrorWithMsg(ctx, errno.ErrReportNotFound, err.Error(), wbReq.RouteUrl)
			return resp, err
		} else {
			resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
			return resp, err
		}
	}

	resp := response.WbSuccessWithData(ctx, result, wbReq.RouteUrl)
	return resp, nil
}

func StressMachineList(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetMachineListParam
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	res, total, err := machine.GetMachineList(ctx, req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.GetMachineListResponse{
		MachineList: res,
		Total:       total,
	}

	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func SendApiResult(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetSendTargetResultReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	apiResult, err := target.GetSendAPIResult(ctx, req.RetID)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccessWithData(ctx, apiResult, wbReq.RouteUrl)
	return resp, nil
}

func SendSceneResult(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetSendSceneResultReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	sceneResult, err := target.GetSendSceneResult(ctx, req.RetID)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	respData := rao.GetSendSceneResultResp{Scenes: sceneResult}

	resp := response.WbSuccessWithData(ctx, respData, wbReq.RouteUrl)
	return resp, nil
}

func DisbandTeamNotice(ctx *gin.Context, wbReq *WebSocketReq) {
	log.Logger.Info("解散团队--wb入参：", *wbReq)
	var req rao.DisbandTeamReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		log.Logger.Info("解散团队--参数错误")
		return
	}

	if userIDs, ok := ClientTeamAndUserMap[req.TeamID]; ok {
		msgType := 1
		for _, userID := range userIDs {
			if websocketLink, ok := ClientLinkList[userID]; ok {
				resp := response.WbSuccess(ctx, wbReq.RouteUrl)
				websocketLink.Mu.Lock()
				defer websocketLink.Mu.Unlock()

				err = websocketLink.Websocket.WriteMessage(msgType, []byte(resp))
				log.Logger.Info("解散团队--循环发送消息", userID)
				if err != nil {
					log.Logger.Info("解散团队--给用户发送消息失败,userID:", userID, "err:", err)
					break
				}
			}
		}
	}
}

func SaveGlobalParam(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.SaveGlobalParamReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	if req.ParamType == 1 {
		tempArr := make([]string, 0, len(req.Cookies))
		for _, v := range req.Cookies {
			isExist := public.StringInSlice(v.Key, tempArr)
			if isExist {
				resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
				return resp, err
			}
			tempArr = append(tempArr, v.Key)
		}
	} else if req.ParamType == 2 {
		tempArr := make([]string, 0, len(req.Headers))
		for _, v := range req.Headers {
			isExist := public.StringInSlice(v.Key, tempArr)
			if isExist {
				resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
				return resp, err
			}
			tempArr = append(tempArr, v.Key)
		}
	} else if req.ParamType == 3 {
		tempArr := make([]string, 0, len(req.Variables))
		for _, v := range req.Variables {
			isExist := public.StringInSlice(v.Key, tempArr)
			if isExist {
				resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
				return resp, err
			}
			tempArr = append(tempArr, v.Key)
		}
	} else if req.ParamType == 4 {

	} else {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
		return resp, err
	}

	err = variable.SaveGlobalParam(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccess(ctx, wbReq.RouteUrl)
	return resp, nil
}

func GetGlobalParam(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetGlobalParamReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	globalParamData, err := variable.GetGlobalParam(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccessWithData(ctx, globalParamData, wbReq.RouteUrl)
	return resp, nil
}

func SaveSceneParam(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.SaveSceneParamReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	if req.ParamType == 1 {
		tempArr := make([]string, 0, len(req.Cookies))
		for _, v := range req.Cookies {
			isExist := public.StringInSlice(v.Key, tempArr)
			if isExist {
				resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
				return resp, err
			}
			tempArr = append(tempArr, v.Key)
		}
	} else if req.ParamType == 2 {
		tempArr := make([]string, 0, len(req.Headers))
		for _, v := range req.Headers {
			isExist := public.StringInSlice(v.Key, tempArr)
			if isExist {
				resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
				return resp, err
			}
			tempArr = append(tempArr, v.Key)
		}
	} else if req.ParamType == 3 {
		tempArr := make([]string, 0, len(req.Variables))
		for _, v := range req.Variables {
			isExist := public.StringInSlice(v.Key, tempArr)
			if isExist {
				resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
				return resp, err
			}
			tempArr = append(tempArr, v.Key)
		}
	} else if req.ParamType == 4 {

	} else {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, "", wbReq.RouteUrl)
		return resp, err
	}

	err = variable.SaveSceneParam(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccess(ctx, wbReq.RouteUrl)
	return resp, nil
}

func GetSceneParam(ctx *gin.Context, wbReq *WebSocketReq) (string, error) {
	var req rao.GetSceneParamReq
	err := json.Unmarshal([]byte(wbReq.Param), &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrParam, err.Error(), wbReq.RouteUrl)
		return resp, err
	}

	sceneParamData, err := variable.GetSceneParam(ctx, &req)
	if err != nil {
		resp := response.WbErrorWithMsg(ctx, errno.ErrMongoFailed, err.Error(), wbReq.RouteUrl)
		return resp, err
	}
	resp := response.WbSuccessWithData(ctx, sceneParamData, wbReq.RouteUrl)
	return resp, nil
}
