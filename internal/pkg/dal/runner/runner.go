package runner

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/errno"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/response"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"
	"time"
)

type RunAPIResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type StopRunnerReq struct {
	TeamID    string   `json:"team_id"`
	PlanID    string   `json:"plan_id"`
	ReportIds []string `json:"report_ids"`
}

func StopSceneCase(ctx *gin.Context, req *rao.StopSceneCaseReq) error {
	// 停止计划的时候，往redis里面写一条数据
	stopSceneCaseKey := consts.StopScenePrefix + req.TeamID + ":" + req.SceneID + ":" + req.SceneCaseID
	_, err := dal.GetRDB().Set(ctx, stopSceneCaseKey, "stop", time.Second*3600).Result()
	if err != nil {
		proof.Errorf("停止场景用例--写入redis数据失败，err:", err)
		response.ErrorWithMsg(ctx, errno.ErrRedisFailed, err.Error())
		return err
	}

	//var ret RunAPIResp
	//_, err := resty.New().R().
	//	SetBody(req).
	//	SetResult(&ret).
	//	Post(conf.Conf.Clients.Runner.StopScene)
	//
	//if err != nil {
	//	return err
	//}
	//
	//if ret.Code != 200 {
	//	return fmt.Errorf("ret code not 200")
	//}

	return nil
}
