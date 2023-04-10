package mao

type ReportData struct {
	//PlanID      string `bson:"plan_id" json:"plan_id"`
	//TeamID      string `bson:"team_id" json:"team_id"`
	//ReportID    string `bson:"report_id" json:"report_id"`
	//Analysis    int64  `bson:"analysis" json:"analysis"`
	//Data        string `bson:"data" json:"data"`
	//Description string `bson:"description" json:"description"`

	PlanID      string `json:"plan_id"`
	TeamID      string `json:"team_id"`
	ReportID    string `json:"report_id"`
	Analysis    int64  `json:"analysis"`
	Data        string `json:"data"`
	Description string `json:"description"`
}
