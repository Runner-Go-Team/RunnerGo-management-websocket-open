package crontab

import (
	"RunnerGo-management/internal/pkg/biz/consts"
	"RunnerGo-management/internal/pkg/biz/log"
	"RunnerGo-management/internal/pkg/dal"
	"golang.org/x/net/context"
	"gorm.io/gen"
	"time"
)

func DeleteOverdueOrder() {
	// 开启定时任务轮询
	for {
		ctx := context.Background()
		tx := dal.GetQuery().Order
		// 组装查询条件
		currentTime := time.Now()
		overdueTimeTemp, _ := time.ParseDuration("-0.5h")
		overdueTime := currentTime.Add(overdueTimeTemp)

		conditions := make([]gen.Condition, 0)
		conditions = append(conditions, tx.PayStatus.Eq(consts.OrderPayStatusNoPay))
		conditions = append(conditions, tx.CreatedAt.Lt(overdueTime))
		// 从数据库当中，查出当前需要执行的定时任务
		overdueOrderList, err := tx.WithContext(ctx).Where(conditions...).Find()
		orderIDs := make([]string, 0, len(overdueOrderList))
		deleteTeamIDs := make([]string, 0, len(overdueOrderList))
		for _, orderInfo := range overdueOrderList {
			orderIDs = append(orderIDs, orderInfo.OrderID)
			if orderInfo.OrderType == consts.OrderTypeCreateNewTeam {
				deleteTeamIDs = append(deleteTeamIDs, orderInfo.TeamID)
			}
		}

		// 删除过期订单
		_, err = tx.WithContext(ctx).Where(tx.OrderID.In(orderIDs...)).Delete()
		if err != nil {
			log.Logger.Info("删除过期订单--定时删除过期订单失败，涉及的订单ID为：", orderIDs)
		}

		// 如果有新建团队订单，则把对应的团队删除掉
		if len(deleteTeamIDs) > 0 {
			teamTable := dal.GetQuery().Team
			_, err = teamTable.WithContext(ctx).Where(teamTable.TeamID.In(deleteTeamIDs...)).Delete()
			if err != nil {
				log.Logger.Info("删除过期订单--定时删除过期订单下的团队失败，涉及的订单ID为：", orderIDs, " 团队ID为:", deleteTeamIDs)
			}
		}

		// 一秒钟循环一次，再循环执行
		time.Sleep(1 * time.Second)
	}
}
