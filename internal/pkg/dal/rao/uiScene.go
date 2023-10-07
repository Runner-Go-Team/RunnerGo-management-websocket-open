package rao

type UIEngineResultDataMsg struct {
	TopicID       string                   `json:"topic"`
	UserID        string                   `json:"user_id"`
	OperatorID    string                   `json:"operator_id"`                  // 操作ID
	SceneID       string                   `json:"scene_id"`                     // 场景
	Sort          int32                    `json:"sort"`                         // 步骤
	ExecTime      float64                  `json:"exec_time"`                    // 执行时间
	Status        string                   `json:"status"`                       // 状态
	RunStatus     int32                    `json:"run_status" bson:"run_status"` // 1:未测 2:成功  3:失败
	Msg           string                   `json:"msg"`
	Screenshot    string                   `json:"screenshot"`
	ScreenshotUrl string                   `json:"screenshot_url"`
	End           bool                     `json:"end"`
	Assertions    []*UIEngineAssertion     `json:"assertions"`
	DataWithdraws []*UIEngineDataWithdraw  `json:"data_withdraws"`
	Withdraws     []*UIEngineDataWithdraw  `json:"withdraws"`
	IsMulti       bool                     `json:"is_multi"`     // 是否展示多条
	IsReport      bool                     `json:"is_report"`    // 是否是报告
	MultiResult   []*UIEngineResultDataMsg `json:"multi_result"` // 多条数据结果
}

type UIEngineAssertion struct {
	Name   string `json:"name"`
	Status bool   `json:"status"` // 状态
	Msg    string `json:"msg"`
}

type UIEngineDataWithdraw struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
