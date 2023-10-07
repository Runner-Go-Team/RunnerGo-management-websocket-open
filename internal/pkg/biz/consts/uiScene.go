package consts

const (
	UISceneRunStatusNo      = 1 // 1:未测 2:成功  3:失败
	UISceneRunStatusSuccess = 2
	UISceneRunStatusFail    = 3

	UIReportStatusIng = 1 // 报告状态1:进行中，2:已完成
	UIReportStatusEnd = 2

	UIEngineCurrentRunPrefix = "UIEngineCurrentRun:" // ui机器对应的运行ID
	UIEngineRunAddrPrefix    = "UIEngineRunAddr:"    // 运行ID对应的机器

	UISceneOptTypeOpenPage        = "open_page"
	UISceneOptTypeClosePage       = "close_page"
	UISceneOptTypeToggleWindow    = "toggle_window"
	UISceneOptTypeForward         = "forward"
	UISceneOptTypeBack            = "back"
	UISceneOptTypeRefresh         = "refresh"
	UISceneOptTypeSetWindowSize   = "set_window_size"
	UISceneOptTypeMouseClicking   = "mouse_clicking"
	UISceneOptTypeMouseScrolling  = "mouse_scrolling"
	UISceneOptTypeMouseMovement   = "mouse_movement"
	UISceneOptTypeMouseDragging   = "mouse_dragging"
	UISceneOptTypeInputOperations = "input_operations"
	UISceneOptTypeWaitEvents      = "wait_events"
	UISceneOptTypeIfCondition     = "if_condition"
	UISceneOptTypeForLoop         = "for_loop"
	UISceneOptTypeWhileLoop       = "while_loop"
	UISceneOptTypeAssert          = "assert"
	UISceneOptTypeDataWithdraw    = "data_withdraw"
)
