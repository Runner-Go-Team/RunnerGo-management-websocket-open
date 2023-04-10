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

func DeleteMongodbData() {

	crontab := cron.New()
	EntryID, err := crontab.AddFunc("* 3 * * *", DeleteDebugData)
	fmt.Println(time.Now(), EntryID, err)

	crontab.Start()
	time.Sleep(time.Minute * 5)

}

func DeleteDebugData() {
	ctx := context.Background()
	// 删除Api调试数据
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectAPIDebug)
	findFilter := bson.D{{}}
	_, err := collection.DeleteOne(ctx, findFilter)
	if err != nil {
		log.Logger.Info("删除操作日志--api_debug日志删除失败，err:", err)
	}
	log.Logger.Info("删除操作日志--api_debug删除成功")

	// 删除场景调试数据
	collection = dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectSceneDebug)
	findFilter = bson.D{{}}
	_, err = collection.DeleteOne(ctx, findFilter)
	if err != nil {
		log.Logger.Info("删除操作日志--scene_debug日志删除失败，err:", err)
	}
	log.Logger.Info("删除操作日志--scene_debug删除成功")
}
