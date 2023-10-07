package packer

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

func TransPlansToRaoPlanList(plans []*model.StressPlan, users []*model.User) []*rao.StressPlan {
	ret := make([]*rao.StressPlan, 0)

	memo := make(map[string]*model.User)
	for _, user := range users {
		memo[user.UserID] = user
	}

	for _, p := range plans {
		ret = append(ret, &rao.StressPlan{
			ID:                p.ID,
			PlanID:            p.PlanID,
			RankID:            p.RankID,
			TeamID:            p.TeamID,
			PlanName:          p.PlanName,
			TaskType:          p.TaskType,
			TaskMode:          p.TaskMode,
			Status:            p.Status,
			CreatedUserName:   memo[p.CreateUserID].Nickname,
			CreatedUserAvatar: memo[p.CreateUserID].Avatar,
			CreatedUserID:     p.CreateUserID,
			Remark:            p.Remark,
			CreatedTimeSec:    p.CreatedAt.Unix(),
			UpdatedTimeSec:    p.UpdatedAt.Unix(),
		})
	}
	return ret
}

func TransTaskToRaoPlan(p *model.StressPlan, t rao.ModeConf, u *model.User) *rao.StressPlan {
	mc := rao.ModeConf{
		RoundNum:         t.RoundNum,
		Concurrency:      t.Concurrency,
		ThresholdValue:   t.ThresholdValue,
		StartConcurrency: t.StartConcurrency,
		Step:             t.Step,
		StepRunTime:      t.StepRunTime,
		MaxConcurrency:   t.MaxConcurrency,
		Duration:         t.Duration,
	}

	return &rao.StressPlan{
		PlanID:            p.PlanID,
		TeamID:            p.TeamID,
		PlanName:          p.PlanName,
		TaskType:          p.TaskType,
		TaskMode:          p.TaskMode,
		Status:            p.Status,
		CreatedUserID:     p.CreateUserID,
		CreatedUserAvatar: u.Avatar,
		CreatedUserName:   u.Nickname,
		Remark:            p.Remark,
		CreatedTimeSec:    p.CreatedAt.Unix(),
		UpdatedTimeSec:    p.UpdatedAt.Unix(),
		ModeConf:          &mc,
	}
}
