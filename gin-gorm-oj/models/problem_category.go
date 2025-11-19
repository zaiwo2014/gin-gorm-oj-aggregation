package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     string         `gorm:"column:problem_id;type:varchar(36)" json:"problem_id"`       //问题 ID
	CategoryId    string         `gorm:"column:category_id;type:varchar(36)" json:"category_id"`     //分类 ID
	CategoryBasic *CategoryBasic `gorm:"foreignKey:ID;references:category_id" json:"category_basic"` //分类信息
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
