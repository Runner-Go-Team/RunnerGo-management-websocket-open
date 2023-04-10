package consts

const (
	OrderTypeCreateNewTeam     = 1 // 新建团队
	OrderTypeVumVersion        = 2 // vum套餐
	OrderTypeUpgradeVumVersion = 3 // 升级团队
	OrderTypeAddUser           = 4 // 增加席位
	OrderTypeAddTime           = 5 // 套餐续期

	ProductTypePerson      = 1 // 个人版
	ProductTypeTeamVersion = 2 // 团队版
	ProductTypeCompany     = 3 // 企业版

	// 团队套餐-月份对应的折扣
	DiscountOneMouth    = 1    // 一个月的折扣--不打折
	DiscountThreeMouth  = 0.95 // 三个月折扣
	DiscountSixMouth    = 0.9  // 六个月折扣
	DiscountTwelveMouth = 0.8  // 12个月折扣

	// 试用期团队过期时间天数
	TrialExpirationDayNum = 30 // 30天

	// 支付状态
	OrderPayStatusNoPay   = 0 // 未支付
	OrderPayStatusSucceed = 1 // 支付成功
	OrderPayStatusFail    = 2 // 支付失败
	OrderPayStatusRefund  = 3 // 退款
)
