package packer

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/go-omnibus/proof"
	"go.mongodb.org/mongo-driver/bson"
)

func TransSaveSceneCaseFlowReqToMaoFlow(req *rao.SaveSceneCaseFlowReq) *mao.SceneCaseFlow {
	nodes, err := bson.Marshal(mao.Node{Nodes: req.Nodes})
	if err != nil {
		proof.Errorf("flow.nodes bson marshal err %w", err)
	}

	edges, err := bson.Marshal(mao.Edge{Edges: req.Edges})
	if err != nil {
		proof.Errorf("flow.edges bson marshal err %w", err)
	}

	return &mao.SceneCaseFlow{
		SceneID:     req.SceneID,
		SceneCaseID: req.SceneCaseID,
		TeamID:      req.TeamID,
		EnvID:       req.EnvID,
		Version:     req.Version,
		Nodes:       nodes,
		Edges:       edges,
		PlanID:      req.PlanID,
	}
}
