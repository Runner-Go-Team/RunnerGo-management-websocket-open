// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserEverydayActivity = "user_everyday_activity"

// UserEverydayActivity mapped from table <user_everyday_activity>
type UserEverydayActivity struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键ID
	UserID    string         `gorm:"column:user_id;not null" json:"user_id"`                                 // 用户ID
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;not null" json:"deleted_at"`                           // 删除时间
}

// TableName UserEverydayActivity's table name
func (*UserEverydayActivity) TableName() string {
	return TableNameUserEverydayActivity
}