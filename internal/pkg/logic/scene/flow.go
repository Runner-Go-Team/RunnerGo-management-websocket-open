package scene

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/errno"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/packer"
)

func SaveFlow(ctx *gin.Context, req *rao.SaveFlowReq) (int, error) {
	flow := packer.TransSaveFlowReqToMaoFlow(req)
	collection := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectFlow)
	err := collection.FindOne(ctx, bson.D{{"scene_id", req.SceneID}}).Err()
	if err == mongo.ErrNoDocuments { // 新建
		_, err := collection.InsertOne(ctx, flow)
		if err != nil {
			return errno.ErrMongoFailed, err
		}
	} else {
		_, err = collection.UpdateOne(ctx, bson.D{
			{"scene_id", req.SceneID},
		}, bson.M{"$set": flow})
		if err != nil {
			return errno.ErrMongoFailed, err
		}

		// 更新场景下所有用例的env_id
		collection2 := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectSceneCaseFlow)
		cur, err := collection2.Find(ctx, bson.D{{"scene_id", req.SceneID}})
		if err == nil {
			var sceneCaseFlow []*mao.SceneCaseFlow
			if err = cur.All(ctx, &sceneCaseFlow); err != nil {
				return errno.ErrMongoFailed, err
			}

			for _, caseFlowInfo := range sceneCaseFlow {
				caseFlowInfo.EnvID = req.EnvID
				_, err = collection2.UpdateOne(ctx, bson.D{
					{"scene_case_id", caseFlowInfo.SceneCaseID},
				}, bson.M{"$set": caseFlowInfo})
			}
		}
	}

	// 保存引入接口和被引入地址的关系
	_ = SaveTargetCiteRelation(ctx, req.Nodes, req.SceneID, req.PlanID, req.TeamID, "", req.Source)
	return errno.Ok, err
}

func SaveTargetCiteRelation(ctx *gin.Context, nodes []rao.Node, sceneID, planID, TeamID, caseID string, source int32) error {
	// 保存引入接口和被引入地址的关系
	collection3 := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectTargetCiteRelation)
	_, err := collection3.DeleteMany(ctx, bson.D{{"scene_id", sceneID}})
	if err != nil {
		return err
	}

	targetCiteRelation := make([]interface{}, 0, len(nodes))
	for _, v := range nodes {
		if v.API.TargetID != "" {
			temp := mao.TargetCiteRelation{
				SceneID:  sceneID,
				TargetID: v.API.TargetID,
				NodeID:   v.ID,
				PlanID:   planID,
				TeamID:   TeamID,
				CaseID:   caseID,
				Source:   source,
			}
			targetCiteRelation = append(targetCiteRelation, temp)
		}
	}
	_, err = collection3.InsertMany(ctx, targetCiteRelation)
	if err != nil {
		return err
	}
	return nil
}
