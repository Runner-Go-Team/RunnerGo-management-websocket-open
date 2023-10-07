package packer

import (
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/mao"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

func TransSaveFlowReqToMaoFlow(req *rao.SaveFlowReq) *mao.Flow {
	nodes, err := bson.Marshal(mao.Node{Nodes: req.Nodes})
	if err != nil {
		log.Logger.Info("flow.nodes bson marshal err %w", err)
	}

	edges, err := bson.Marshal(mao.Edge{Edges: req.Edges})
	if err != nil {
		log.Logger.Info("flow.edges bson marshal err %w", err)
	}

	prepositions, err := bson.Marshal(mao.Preposition{Prepositions: req.Prepositions})
	if err != nil {
		log.Logger.Info("flow.prepositions bson marshal err %w", err)
	}

	return &mao.Flow{
		SceneID:      req.SceneID,
		TeamID:       req.TeamID,
		EnvID:        req.EnvID,
		Version:      req.Version,
		Nodes:        nodes,
		Edges:        edges,
		Prepositions: prepositions,
		PlanID:       req.PlanID,
	}
}
