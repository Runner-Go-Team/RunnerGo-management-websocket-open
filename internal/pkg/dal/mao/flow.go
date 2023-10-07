package mao

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
)

type Flow struct {
	SceneID      string   `bson:"scene_id"`
	TeamID       string   `bson:"team_id"`
	EnvID        int64    `bson:"env_id"`
	Version      int32    `bson:"version"`
	Nodes        bson.Raw `bson:"nodes"`
	Edges        bson.Raw `bson:"edges"`
	Prepositions bson.Raw `bson:"prepositions,omitempty"`
	PlanID       string   `bson:"plan_id"`
}

type Preposition struct {
	Prepositions []rao.Node `bson:"prepositions"`
}

type Node struct {
	Nodes []rao.Node `bson:"nodes"`
}

type Edge struct {
	Edges []rao.Edge `bson:"edges"`
}
