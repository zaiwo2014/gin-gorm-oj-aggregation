package models

import (
	"time"

	"gorm.io/gorm"
)

// CategoryBasic 分类基础信息模型
type CategoryBasic struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Identity  string         `gorm:"column:identity;type:varchar(36)" json:"identity"`
	Name      string         `gorm:"column:name;type:varchar(100)" json:"name"`
	ParentID  int            `gorm:"column:parent_id;type:int(11);default:0" json:"parent_id"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (table *CategoryBasic) TableName() string {
	return "category_basic"
}
