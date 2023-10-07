package mao

type TargetCiteRelation struct {
	SceneID  string `bson:"scene_id"`
	CaseID   string `bson:"case_id"`
	TargetID string `bson:"target_id"`
	NodeID   string `bson:"node_id"`
	PlanID   string `bson:"plan_id"`
	TeamID   string `bson:"team_id"`
	Source   int32  `bson:"source"`
}
