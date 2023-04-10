// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVariableImport = "variable_import"

// VariableImport mapped from table <variable_import>
type VariableImport struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TeamID     int64          `gorm:"column:team_id" json:"team_id"`   // 团队id
	SceneID    int64          `gorm:"column:scene_id" json:"scene_id"` // 场景id
	Name       string         `gorm:"column:name;not null" json:"name"`
	URL        string         `gorm:"column:url;not null" json:"url"`
	UploaderID int64          `gorm:"column:uploader_id;not null" json:"uploader_id"` // 上传人id
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName VariableImport's table name
func (*VariableImport) TableName() string {
	return TableNameVariableImport
}
