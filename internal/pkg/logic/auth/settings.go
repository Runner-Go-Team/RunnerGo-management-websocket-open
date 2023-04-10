package auth

import (
	"RunnerGo-management/internal/pkg/dal"
	"RunnerGo-management/internal/pkg/dal/model"
	"RunnerGo-management/internal/pkg/dal/query"
	"RunnerGo-management/internal/pkg/dal/rao"
	"RunnerGo-management/internal/pkg/packer"
	"context"
)

func SetUserSettings(ctx context.Context, userID string, settings *rao.UserSettings) error {
	currentTeamID := settings.CurrentTeamID
	//currentDateTime := time.Now()

	//if currentTeamID != "" {
	//	teamTable := query.Use(dal.DB()).Team
	//	_, teamInfoErr := teamTable.WithContext(ctx).Where(teamTable.TeamID.Eq(currentTeamID),
	//		teamTable.VipExpirationDate.Gt(currentDateTime), teamTable.IsUsable.Eq(consts.TeamCanUse)).First()
	//
	//	if teamInfoErr != nil {
	//		return fmt.Errorf("团队已过期")
	//	}
	//}

	tx := query.Use(dal.DB()).Setting
	_, err := tx.WithContext(ctx).Where(tx.UserID.Eq(userID)).UpdateColumnSimple(tx.TeamID.Value(currentTeamID))
	if err != nil {
		return err
	}

	return nil
}

func GetUserSettings(ctx context.Context, userID string) (*rao.GetUserSettingsResp, error) {

	tx := query.Use(dal.DB()).Setting
	settingInfo, err := tx.WithContext(ctx).Where(tx.UserID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}

	userInfo := new(model.User)

	// 查询当前用户在默认团队的角色
	userTeamTable := dal.GetQuery().UserTeam
	utInfo, err := userTeamTable.WithContext(ctx).Where(userTeamTable.TeamID.Eq(settingInfo.TeamID),
		userTeamTable.UserID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}

	// 查询用户信息
	userTable := query.Use(dal.DB()).User
	userInfo, err = userTable.WithContext(ctx).Where(userTable.UserID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}

	return packer.TransUserSettingsToRaoUserSettings(settingInfo, utInfo, userInfo), nil
}

// GetAvailTeamID 获取有效的团队ID
func GetAvailTeamID(ctx context.Context, userID string) (string, error) {

	//获取用户最后一次使用的团队
	tx := query.Use(dal.DB()).Setting
	s, err := tx.WithContext(ctx).Where(tx.UserID.Eq(userID)).First()
	if err != nil {
		return "", err
	}
	lastOperationTeamID := s.TeamID
	//fmt.Println("lastOperationTeamID ==== ", lastOperationTeamID)

	////获取未过期的付费团队
	//userTeamTable := query.Use(dal.DB()).UserTeam
	//userTeamList, userTeamListErr := userTeamTable.WithContext(ctx).Select(userTeamTable.TeamID).Where(userTeamTable.UserID.Eq(userID)).Find()
	//if userTeamListErr != nil || len(userTeamList) == 0 {
	//	return "", userTeamListErr
	//}
	//
	////获取用户所有的团队ID 然后到团队列表查询团队的详情信息
	//var teamID = make([]string, 0, len(userTeamList))
	//for _, userTeamValue := range userTeamList {
	//	teamID = append(teamID, userTeamValue.TeamID)
	//}
	//currentDateTime := time.Now()
	//
	//teamTable := query.Use(dal.DB()).Team
	//teamList, teamListErr := teamTable.WithContext(ctx).Where(teamTable.TeamID.In(teamID...), teamTable.IsUsable.Eq(consts.TeamCanUse)).Find()
	//if teamListErr != nil || len(teamList) == 0 {
	//	return "", teamListErr
	//}
	//
	////获取到付费团队ID
	//vipTeam := make([]model.Team, 0, len(teamList))
	//vipTeamID := make([]string, 0, len(teamList))
	//for _, teamListVal := range teamList {
	//	if currentDateTime.Before(teamListVal.VipExpirationDate) {
	//		vipTeam = append(vipTeam, *teamListVal)
	//		vipTeamID = append(vipTeamID, teamListVal.TeamID)
	//		//如果用户最后一次使用的团队 现在还是付费团队 则直接跳转进去
	//		if lastOperationTeamID == teamListVal.TeamID {
	//			return lastOperationTeamID, nil
	//		}
	//	}
	//}
	//如果当前没有付费团队 则跳转到付费界面
	//if len(vipTeam) == 0 {
	//	return "", nil
	//}

	//如果有付费的团队 则随机取出一个进入
	//availTeamID := vipTeamID[0]

	return lastOperationTeamID, nil
}
