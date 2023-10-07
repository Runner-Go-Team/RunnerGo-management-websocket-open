package uiReport

import (
	"context"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/query"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

func HandleReportEnd(ctx context.Context, resultData *rao.UIEngineResultDataMsg) error {
	// 查询是否存在报告
	pr := query.Use(dal.DB()).UIPlanReport
	planReport, _ := pr.WithContext(ctx).Where(pr.ReportID.Eq(resultData.TopicID)).First()
	if planReport == nil {
		return nil
	}

	var upDataPlanReport = make(map[string]interface{}, 0)
	upDataPlanReport["run_duration_time"] = int32(resultData.ExecTime)
	upDataPlanReport["status"] = consts.UIReportStatusEnd

	if err := query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		if _, err := tx.UIPlanReport.WithContext(ctx).Where(
			tx.UIPlanReport.ReportID.Eq(resultData.TopicID),
		).Updates(upDataPlanReport); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
