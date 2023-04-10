package mao

type ReportTask struct {
	TeamID   string    `bson:"team_id" json:"team_id"`
	PlanID   string    `bson:"plan_id" json:"plan_id"`
	PlanName string    `bson:"plan_name" json:"plan_name"`
	ReportID string    `bson:"report_id" json:"report_id"`
	TaskType int32     `bson:"task_type" json:"task_type"`
	TaskMode int32     `bson:"task_mode" json:"task_mode"`
	ModeConf *ModeConf `bson:"mode_conf" json:"mode_conf"`
}
