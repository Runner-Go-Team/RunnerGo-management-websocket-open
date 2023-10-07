package autoPlan

import (
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"time"
)

func GetAutoPlanList(ctx *gin.Context, req *rao.GetAutoPlanListReq) ([]*rao.AutoPlanDetailResp, int64, error) {
	// 查询数据库
	tx := dal.GetQuery().AutoPlan
	// 查询数据库
	limit := req.Size
	offset := (req.Page - 1) * req.Size
	sort := make([]field.Expr, 0, 6)
	if req.Sort == 0 { // 默认排序
		sort = append(sort, tx.RankID.Desc())
	} else if req.Sort == 2 { // 创建时间倒序
		sort = append(sort, tx.CreatedAt.Desc())
	} else if req.Sort == 1 { // 创建时间升序
		sort = append(sort, tx.CreatedAt)
	} else if req.Sort == 3 { // 最后修改时间升序
		sort = append(sort, tx.UpdatedAt)
	} else { // 最后修改时间倒序
		sort = append(sort, tx.UpdatedAt.Desc())
	}

	conditions := make([]gen.Condition, 0)
	conditions = append(conditions, tx.TeamID.Eq(req.TeamId))

	if req.PlanName != "" {
		conditions = append(conditions, tx.PlanName.Like(fmt.Sprintf("%%%s%%", req.PlanName)))
		// 先查询出来用户id
		userTable := dal.GetQuery().User
		userList, err := userTable.WithContext(ctx).Where(userTable.Nickname.Like(fmt.Sprintf("%%%s%%", req.PlanName))).Find()
		if err == nil {
			tempUserIDs := make([]string, 0, len(userList))
			for _, userInfo := range userList {
				tempUserIDs = append(tempUserIDs, userInfo.UserID)
			}

			// 查询属于当前团队的用户
			userTeamTable := dal.GetQuery().UserTeam
			userTeamList, err := userTeamTable.WithContext(ctx).Where(userTeamTable.TeamID.Eq(req.TeamId),
				userTeamTable.UserID.In(tempUserIDs...)).Find()
			if err == nil {
				userIDs := make([]string, 0, len(userList))
				for _, vutInfo := range userTeamList {
					userIDs = append(userIDs, vutInfo.UserID)
				}
				if len(userIDs) > 0 {
					conditions[1] = tx.RunUserID.In(userIDs...)
				}
			}
		}
	}

	if req.TaskType != 0 {
		conditions = append(conditions, tx.TaskType.Eq(req.TaskType))
	}

	if req.Status != 0 {
		conditions = append(conditions, tx.Status.Eq(req.Status))
	}

	if (req.StartTimeSec != 0 && req.EndTimeSec != 0) && (req.EndTimeSec > req.StartTimeSec) {
		startTime := time.Unix(req.StartTimeSec, 0)
		endTime := time.Unix(req.EndTimeSec, 0)
		conditions = append(conditions, tx.CreatedAt.Between(startTime, endTime))
	}

	list, total, err := tx.WithContext(ctx).Where(conditions...).Order(sort...).FindByPage(offset, limit)
	if err != nil {
		log.Logger.Info("自动化计划列表--获取列表失败，err:", err)
		return nil, 0, err
	}

	// 获取所有操作人id
	runUserIDs := make([]string, 0, len(list))
	for _, detail := range list {
		runUserIDs = append(runUserIDs, detail.RunUserID)
	}

	userTable := dal.GetQuery().User
	userList, err := userTable.WithContext(ctx).Where(userTable.UserID.In(runUserIDs...)).Find()
	if err != nil {
		return nil, 0, err
	}
	// 用户id和名称映射
	userMap := make(map[string]string, len(userList))
	for _, userValue := range userList {
		userMap[userValue.UserID] = userValue.Nickname
	}

	res := make([]*rao.AutoPlanDetailResp, 0, len(list))
	for _, detail := range list {
		detailTmp := &rao.AutoPlanDetailResp{
			RankID:    detail.RankID,
			PlanID:    detail.PlanID,
			TeamID:    detail.TeamID,
			PlanName:  detail.PlanName,
			TaskType:  detail.TaskType,
			CreatedAt: detail.CreatedAt.Unix(),
			UpdatedAt: detail.UpdatedAt.Unix(),
			Status:    detail.Status,
			Remark:    detail.Remark,
			UserName:  userMap[detail.RunUserID],
		}
		res = append(res, detailTmp)
	}
	return res, total, nil
}

func GetAutoPlanDetail(ctx *gin.Context, req *rao.GetAutoPlanDetailReq) (*rao.GetAutoPlanDetailResp, error) {
	// 获取计划详情
	tx := dal.GetQuery().AutoPlan
	planInfo, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(req.TeamID)).Where(tx.PlanID.Eq(req.PlanID)).First()
	if err != nil {
		return nil, err
	}

	// 查询创建人
	tableUser := dal.GetQuery().User
	userInfo, err := tableUser.WithContext(ctx).Where(tableUser.UserID.Eq(planInfo.CreateUserID)).First()
	if err != nil {
		return nil, err
	}

	res := &rao.GetAutoPlanDetailResp{
		PlanID:    planInfo.PlanID,
		TeamID:    planInfo.TeamID,
		PlanName:  planInfo.PlanName,
		CreatedAt: planInfo.CreatedAt.Unix(),
		UpdatedAt: planInfo.UpdatedAt.Unix(),
		Remark:    planInfo.Remark,
		Status:    planInfo.Status,
		UserName:  userInfo.Nickname,
		Avatar:    userInfo.Avatar,
	}
	return res, nil
}

func GetAutoPlanReportList(ctx *gin.Context, req *rao.GetAutoPlanReportListReq) ([]*rao.GetAutoPlanReportList, int64, error) {
	// 查询数据库
	tx := dal.GetQuery().AutoPlanReport
	// 查询数据库
	limit := req.Size
	offset := (req.Page - 1) * req.Size
	sort := make([]field.Expr, 0, 6)
	if req.Sort == 0 { // 默认排序
		sort = append(sort, tx.RankID.Desc())
	} else if req.Sort == 2 { // 创建时间倒序
		sort = append(sort, tx.CreatedAt.Desc())
	} else if req.Sort == 1 { // 创建时间升序
		sort = append(sort, tx.CreatedAt)
	}

	conditions := make([]gen.Condition, 0)
	conditions = append(conditions, tx.TeamID.Eq(req.TeamId))
	if req.PlanName != "" {
		conditions = append(conditions, tx.PlanName.Like(fmt.Sprintf("%%%s%%", req.PlanName)))
		// 先查询出来用户id
		userTable := dal.GetQuery().User
		userList, err := userTable.WithContext(ctx).Where(userTable.Nickname.Like(fmt.Sprintf("%%%s%%", req.PlanName))).Find()
		if err == nil && len(userList) > 0 {
			tempUserIDs := make([]string, 0, len(userList))
			for _, userInfo := range userList {
				tempUserIDs = append(tempUserIDs, userInfo.UserID)
			}

			// 查询属于当前团队的用户
			userTeamTable := dal.GetQuery().UserTeam
			userTeamList, err := userTeamTable.WithContext(ctx).Where(userTeamTable.TeamID.Eq(req.TeamId),
				userTeamTable.UserID.In(tempUserIDs...)).Find()
			if err == nil && len(userTeamList) > 0 {
				userIDs := make([]string, 0, len(userList))
				for _, vutInfo := range userTeamList {
					userIDs = append(userIDs, vutInfo.UserID)
				}
				if len(userIDs) > 0 {
					conditions[1] = tx.RunUserID.In(userIDs...)
				}
			}
		}
	}

	if req.TaskType != 0 {
		conditions = append(conditions, tx.TaskType.Eq(req.TaskType))
	}

	if req.TaskType != 0 {
		conditions = append(conditions, tx.TaskType.Eq(req.TaskType))
	}

	if req.Status != 0 {
		conditions = append(conditions, tx.Status.Eq(req.Status))
	}

	if (req.StartTimeSec != 0 && req.EndTimeSec != 0) && (req.EndTimeSec > req.StartTimeSec) {
		startTime := time.Unix(req.StartTimeSec, 0)
		endTime := time.Unix(req.EndTimeSec, 0)
		conditions = append(conditions, tx.CreatedAt.Between(startTime, endTime))
	}

	list, total, err := tx.WithContext(ctx).Where(conditions...).Order(sort...).FindByPage(offset, limit)
	if err != nil {
		log.Logger.Info("自动化计划报告列表--获取列表失败，err:", err)
		return nil, 0, err
	}

	// 获取所有创建人id
	createUserIDs := make([]string, 0, len(list))
	for _, detail := range list {
		createUserIDs = append(createUserIDs, detail.RunUserID)
	}

	userTable := dal.GetQuery().User
	userList, err := userTable.WithContext(ctx).Where(userTable.UserID.In(createUserIDs...)).Find()
	if err != nil {
		return nil, 0, err
	}
	// 用户id和名称映射
	userMap := make(map[string]string, len(userList))
	for _, userValue := range userList {
		userMap[userValue.UserID] = userValue.Nickname
	}

	res := make([]*rao.GetAutoPlanReportList, 0, len(list))
	for _, detail := range list {
		detailTmp := &rao.GetAutoPlanReportList{
			RankID:           detail.RankID,
			ReportID:         detail.ReportID,
			ReportName:       detail.ReportName,
			PlanID:           detail.PlanID,
			TeamID:           detail.TeamID,
			PlanName:         detail.PlanName,
			TaskType:         detail.TaskType,
			TaskMode:         detail.TaskMode,
			SceneRunOrder:    detail.SceneRunOrder,
			TestCaseRunOrder: detail.TestCaseRunOrder,
			StartTimeSec:     detail.CreatedAt.Unix(),
			EndTimeSec:       detail.UpdatedAt.Unix(),
			Status:           detail.Status,
			Remark:           detail.Remark,
			RunUserName:      userMap[detail.RunUserID],
		}
		res = append(res, detailTmp)
	}
	return res, total, nil
}

type TestCaseResult struct {
	CaseName   string    `json:"case_name" bson:"case_name"`
	SucceedNum int64     `json:"succeed_num" bson:"succeed_num"`
	TotalNum   int64     `json:"total_num" bson:"total_num"`
	ApiList    []ApiList `json:"api_list" bson:"api_list"`
}
type ApiList struct {
	EventID        string         `json:"event_id" bson:"event_id"`
	ApiName        string         `json:"api_name" bson:"api_name"`
	Method         string         `json:"method" bson:"method"`
	Url            string         `json:"url" bson:"url"`
	Status         string         `json:"status" bson:"status"`
	ResponseBytes  float64        `json:"response_bytes" bson:"response_bytes"`
	RequestTime    int64          `json:"request_time" bson:"request_time"`
	RequestCode    int32          `json:"request_code" bson:"request_code"`
	RequestHeader  string         `json:"request_header" bson:"request_header"`
	RequestBody    string         `json:"request_body" bson:"request_body"`
	ResponseHeader string         `json:"response_header" bson:"response_header"`
	ResponseBody   string         `json:"response_body" bson:"response_body"`
	AssertionMsg   []AssertionMsg `json:"assertion_msg" bson:"assertion_msg"`
}

type AssertionMsg struct {
	Type      string `json:"type"`
	Code      int64  `json:"code" bson:"code"`
	IsSucceed bool   `json:"isSucceed" bson:"isSucceed"`
	Msg       string `json:"msg" bson:"msg"`
}

// SceneResult 场景结果
type SceneResult struct {
	SceneID      string `json:"scene_id" bson:"scene_id"`
	SceneName    string `json:"scene_name" bson:"scene_name"`
	CaseFailNum  int    `json:"case_fail_num" bson:"case_fail_num"`
	CaseTotalNum int    `json:"case_total_num" bson:"case_total_num"`
	State        int    `json:"state" bson:"state"` // 1-成功，2-失败
}

// GetReportDetailResp 获取报告详情返回值
type GetReportDetailResp struct {
	PlanName             string                      `json:"plan_name" bson:"plan_name"`
	Avatar               string                      `json:"avatar" bson:"avatar"`
	Nickname             string                      `json:"nickname" bson:"nickname"`
	Remark               string                      `json:"remark" bson:"remark"`
	TaskMode             int32                       `json:"task_mode" bson:"task_mode"`
	SceneRunOrder        int32                       `json:"scene_run_order" bson:"scene_run_order"`
	TestCaseRunOrder     int32                       `json:"test_case_run_order" bson:"test_case_run_order"`
	ReportStatus         int32                       `json:"report_status" bson:"report_status"`
	ReportStartTime      int64                       `json:"report_start_time" bson:"report_start_time"`
	ReportEndTime        int64                       `json:"report_end_time" bson:"report_end_time"`
	ReportRunTime        int64                       `json:"report_run_time" bson:"report_run_time"`
	SceneBaseInfo        SceneBaseInfo               `json:"scene_base_info" bson:"scene_base_info"`
	CaseBaseInfo         CaseBaseInfo                `json:"case_base_info" bson:"case_base_info"`
	ApiBaseInfo          ApiBaseInfo                 `json:"api_base_info" bson:"api_base_info"`
	AssertionBaseInfo    AssertionBaseInfo           `json:"assertion_base_info" bson:"assertion_base_info"`
	SceneResult          []SceneResult               `json:"scene_result" bson:"scene_result"`
	SceneIDCaseResultMap map[string][]TestCaseResult `json:"scene_id_case_result_map" bson:"scene_id_case_result_map"`
}
type AssertionBaseInfo struct {
	AssertionTotalNum int64 `json:"assertion_total_num" bson:"assertion_total_num"`
	SucceedNum        int64 `json:"succeed_num" bson:"succeed_num"`
	FailNum           int64 `json:"fail_num" bson:"fail_num"`
}

type ApiBaseInfo struct {
	ApiTotalNum int64 `json:"api_total_num" bson:"api_total_num"`
	SucceedNum  int64 `json:"succeed_num" bson:"succeed_num"`
	FailNum     int64 `json:"fail_num" bson:"fail_num"`
	NotTestNum  int64 `json:"not_test_num" bson:"not_test_num"`
}

type CaseBaseInfo struct {
	CaseTotalNum int64 `json:"case_total_num" bson:"case_total_num"`
	SucceedNum   int64 `json:"succeed_num" bson:"succeed_num"`
	FailNum      int64 `json:"fail_num" bson:"fail_num"`
}

type SceneBaseInfo struct {
	SceneTotalNum int64 `json:"scene_total_num" bson:"scene_total_num"`
}
