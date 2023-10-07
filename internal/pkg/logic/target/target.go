package target

import (
	"context"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/packer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSendSceneResult(ctx context.Context, retID string) ([]*rao.SceneDebug, error) {
	cur, err := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectSceneDebug).
		Find(ctx, bson.D{{"uuid", retID}})
	if err != nil {
		return nil, err
	}

	var sds []*mao.SceneDebug
	if err := cur.All(ctx, &sds); err != nil {
		return nil, err
	}

	if len(sds) == 0 {
		return nil, nil
	}

	return packer.TransMaoSceneDebugsToRaoSceneDebugs(sds), nil

}

func GetSendAPIResult(ctx context.Context, retID string) (*rao.APIDebug, error) {
	var ad mao.APIDebug
	err := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectAPIDebug).
		FindOne(ctx, bson.D{{"uuid", retID}}).Decode(&ad)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return packer.TransMaoAPIDebugToRaoAPIDebug(&ad), nil
}
