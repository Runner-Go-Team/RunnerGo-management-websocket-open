package plan

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"gorm.io/gorm"
	"time"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/query"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/packer"
)

func ListByStatus(ctx context.Context, teamID string) (int, error) {
	runPlanNum := 0
	tx := query.Use(dal.DB()).StressPlan
	stressPlanCount, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.Status.Eq(consts.PlanStatusUnderway)).Count()
	if err != nil {
		return 0, err
	}
	tx2 := dal.GetQuery().AutoPlan
	autoPlanCount, err := tx2.WithContext(ctx).Where(tx2.TeamID.Eq(teamID), tx2.Status.Eq(consts.PlanStatusUnderway)).Count()
	if err != nil {
		return 0, err
	}
	runPlanNum = int(stressPlanCount) + int(autoPlanCount)
	return runPlanNum, nil
}

func ListByTeamID(ctx context.Context, req *rao.ListPlansReq) ([]*rao.StressPlan, int64, error) {
	tx := query.Use(dal.DB()).StressPlan
	conditions := make([]gen.Condition, 0)
	conditions = append(conditions, tx.TeamID.Eq(req.TeamID))

	conditions2 := make([]gen.Condition, 0)

	if req.Keyword != "" {
		conditions = append(conditions, tx.PlanName.Like(fmt.Sprintf("%%%s%%", req.Keyword)))

		u := query.Use(dal.DB()).User
		users, err := u.WithContext(ctx).Where(u.Nickname.Like(fmt.Sprintf("%%%s%%", req.Keyword))).Find()
		if err != nil {
			return nil, 0, err
		}

		userIds := make([]string, 0, len(users))
		if len(users) > 0 {
			for _, userInfo := range users {
				userIds = append(userIds, userInfo.UserID)
			}
			conditions2 = append(conditions, tx.RunUserID.In(userIds...))
		}
	}

	if req.StartTimeSec != 0 && req.EndTimeSec != 0 {
		startTime := time.Unix(req.StartTimeSec, 0)
		endTime := time.Unix(req.EndTimeSec, 0)
		conditions = append(conditions, tx.CreatedAt.Between(startTime, endTime))
	}

	if req.TaskType > 0 {
		conditions = append(conditions, tx.TaskType.Eq(req.TaskType))
	}

	if req.TaskMode > 0 {
		conditions = append(conditions, tx.TaskMode.Eq(req.TaskMode))
	}

	if req.Status > 0 {
		conditions = append(conditions, tx.Status.Eq(req.Status))
	}

	sort := make([]field.Expr, 0)
	if req.Sort == 0 { // 默认排序
		sort = append(sort, tx.RankID.Desc())
	}
	if req.Sort == 1 { // 创建时间倒序
		sort = append(sort, tx.CreatedAt.Desc())
	}
	if req.Sort == 2 { // 创建时间正序
		sort = append(sort, tx.CreatedAt)
	}
	if req.Sort == 3 { // 修改时间倒序
		sort = append(sort, tx.UpdatedAt.Desc())
	}
	if req.Sort == 4 { // 修改时间正序
		sort = append(sort, tx.UpdatedAt)
	}

	ret := make([]*model.StressPlan, 0, req.Size)
	var cnt int64 = 0
	var err error

	offset := (req.Page - 1) * req.Size
	limit := req.Size
	if len(conditions2) > 0 {
		ret, cnt, err = tx.WithContext(ctx).Where(conditions...).Or(conditions2...).Order(sort...).FindByPage(offset, limit)
	} else {
		ret, cnt, err = tx.WithContext(ctx).Where(conditions...).Order(sort...).FindByPage(offset, limit)
	}
	if err != nil {
		return nil, 0, err
	}

	var userIDs []string
	for _, r := range ret {
		userIDs = append(userIDs, r.CreateUserID)
	}

	u := query.Use(dal.DB()).User
	users, err := u.WithContext(ctx).Where(u.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, 0, err
	}

	return packer.TransPlansToRaoPlanList(ret, users), cnt, nil
}

func GetByPlanID(ctx context.Context, teamID string, planID string) (*rao.StressPlan, error) {
	tx := dal.GetQuery().StressPlan
	planInfo, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.PlanID.Eq(planID)).First()
	if err != nil {
		return nil, err
	}

	// 查询用户信息
	u := query.Use(dal.DB()).User
	user, err := u.WithContext(ctx).Where(u.UserID.Eq(planInfo.CreateUserID)).First()
	if err != nil {
		return nil, err
	}

	// 查询配置信息
	taskConfTable := dal.GetQuery().StressPlanTaskConf
	taskConfInfo, err := taskConfTable.WithContext(ctx).Where(taskConfTable.TeamID.Eq(teamID), taskConfTable.PlanID.Eq(planID)).Order(taskConfTable.SceneID).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	var taskConf rao.ModeConf
	if err == nil {
		err := json.Unmarshal([]byte(taskConfInfo.ModeConf), &taskConf)
		if err != nil {
			log.Logger.Info("性能计划--任务配置数据解析失败，配置为：", taskConfInfo.ModeConf)
		}
	}

	return packer.TransTaskToRaoPlan(planInfo, taskConf, user), nil
}
