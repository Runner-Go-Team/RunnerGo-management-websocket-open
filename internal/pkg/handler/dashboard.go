package handler

import (
	"RunnerGo-management/internal/pkg/biz/errno"
	"RunnerGo-management/internal/pkg/biz/jwt"
	"RunnerGo-management/internal/pkg/biz/response"
	"RunnerGo-management/internal/pkg/dal/rao"
	"RunnerGo-management/internal/pkg/logic/homePage"
	"RunnerGo-management/internal/pkg/logic/operation"
	"RunnerGo-management/internal/pkg/logic/plan"
	"RunnerGo-management/internal/pkg/logic/report"
	"RunnerGo-management/internal/pkg/logic/target"
	"RunnerGo-management/internal/pkg/logic/user"

	"github.com/gin-gonic/gin"
)

// DashboardDefault 首页控制台
func DashboardDefault(ctx *gin.Context) {
	var req rao.DashboardDefaultReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	u, err := user.FirstByUserID(ctx, jwt.GetUserIDByCtx(ctx), req.TeamID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	operations, _, err := operation.List(ctx, req.TeamID, 5, 0)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	apiCnt, err := target.APICountByTeamID(ctx, req.TeamID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	sceneCnt, err := target.SceneCountByTeamID(ctx, req.TeamID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	planCnt, err := plan.CountByTeamID(ctx, req.TeamID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	reportCnt, err := report.CountByTeamID(ctx, req.TeamID)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.DashboardDefaultResp{
		User:       u,
		Operations: operations,
		PlanNum:    planCnt,
		SceneNum:   sceneCnt,
		ReportNum:  reportCnt,
		APINum:     apiCnt,
		Mobile:     u.Mobile,
	})
	return
}

// HomePage 新的首页
func HomePage(ctx *gin.Context) {
	var req rao.HomePageReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}
	homePageData, err := homePage.HomePage(ctx, &req)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, homePageData)
	return
}
