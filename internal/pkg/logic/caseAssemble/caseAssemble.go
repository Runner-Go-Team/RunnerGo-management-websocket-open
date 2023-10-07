package caseAssemble

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/logic/scene"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/packer"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveSceneCaseFlow(ctx *gin.Context, req *rao.SaveSceneCaseFlowReq) error {
	flow := packer.TransSaveSceneCaseFlowReqToMaoFlow(req)
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectSceneCaseFlow)

	err := collection.FindOne(ctx, bson.D{{"scene_case_id", req.SceneCaseID}}).Err()
	if err == mongo.ErrNoDocuments { // 新建
		_, err = collection.InsertOne(ctx, flow)
	} else {
		_, err = collection.UpdateOne(ctx, bson.D{
			{"scene_case_id", req.SceneCaseID},
		}, bson.M{"$set": flow})
	}

	// 保存引入接口和被引入地址的关系
	_ = scene.SaveTargetCiteRelation(ctx, req.Nodes, req.SceneID, req.PlanID, req.TeamID, req.SceneCaseID, req.Source)

	return err
}
