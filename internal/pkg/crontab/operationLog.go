package crontab

import (
	"RunnerGo-management/internal/pkg/biz/consts"
	"RunnerGo-management/internal/pkg/biz/log"
	"RunnerGo-management/internal/pkg/dal"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func DeleteOperationLogBeforeSevenDay() {

	crontab := cron.New()
	EntryID, err := crontab.AddFunc("* 3 * * *", DeleteOperationLog)
	fmt.Println(time.Now(), EntryID, err)

	crontab.Start()
	time.Sleep(time.Minute * 5)

}

func DeleteOperationLog() {
	nowTime := time.Now().Unix()
	beforeSevenDay := nowTime - (3600 * 24 * 7)

	ctx := context.Background()
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectOperationLog)
	findFilter := bson.D{{"created_time_sec", bson.D{{"$lt", beforeSevenDay}}}}
	_, err := collection.DeleteMany(ctx, findFilter)
	if err != nil {
		log.Logger.Info("删除操作日志--删除失败，err:", err)
	}
	log.Logger.Info("删除操作日志--删除成功")
}
