// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePlan = "plan"

// Plan mapped from table <plan>
type Plan struct {
	ID                 int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键
	TeamID             int64          `gorm:"column:team_id;not null" json:"team_id"`            // 团队ID
	Rank               int64          `gorm:"column:rank;not null" json:"rank"`                  // 团队内份数
	Name               string         `gorm:"column:name;not null" json:"name"`                  // 计划名称
	TaskType           int32          `gorm:"column:task_type" json:"task_type"`                 // 计划类型{1:普通任务,2:定时任务}
	Mode               int32          `gorm:"column:mode" json:"mode"`                           // 压测类型 1 // 并发模式，  2 // 阶梯模式，  3 // 错误率模式，  4 // 响应时间模式，  5 //每秒请求数模式，  6 //每秒事务数模式，
	Status             int32          `gorm:"column:status;not null" json:"status"`              // 计划状态1:未开始,2:进行中
	CreateUserIdentify string         `gorm:"column:create_user_identify;not null" json:"create_user_identify"`
	RunUserIdentify    string         `gorm:"column:run_user_identify;not null" json:"run_user_identify"`
	CreateUserID       int64          `gorm:"column:create_user_id;not null" json:"create_user_id"` // 创建人id
	RunUserID          int64          `gorm:"column:run_user_id;not null" json:"run_user_id"`       // 运行人id
	Remark             string         `gorm:"column:remark" json:"remark"`                          // 备注
	CronExpr           string         `gorm:"column:cron_expr" json:"cron_expr"`                    // 定时任务表达式
	CreatedAt          time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Plan's table name
func (*Plan) TableName() string {
	return TableNamePlan
}
