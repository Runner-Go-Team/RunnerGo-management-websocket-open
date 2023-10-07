// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUISceneSync = "ui_scene_sync"

// UISceneSync mapped from table <ui_scene_sync>
type UISceneSync struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键ID
	SceneID       string         `gorm:"column:scene_id;not null" json:"scene_id"`                               // 场景ID
	SourceSceneID string         `gorm:"column:source_scene_id;not null" json:"source_scene_id"`                 // 引用场景ID
	TeamID        string         `gorm:"column:team_id;not null" json:"team_id"`                                 // 团队id
	SyncMode      int32          `gorm:"column:sync_mode;not null" json:"sync_mode"`                             // 状态：1-实时，2-手动,已场景为准   3-手动,已计划为准
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 修改时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                    // 删除时间
}

// TableName UISceneSync's table name
func (*UISceneSync) TableName() string {
	return TableNameUISceneSync
}
