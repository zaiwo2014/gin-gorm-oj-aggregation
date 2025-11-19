package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;size:36" json:"identity"`
	Title             string             `gorm:"column:title;size:255" json:"title"`    // 问题的题目
	Content           string             `gorm:"column:content" json:"content"`         // 问题的正文描述
	MaxRuntime        int                `gorm:"column:max_runtime" json:"max_runtime"` // 最大的运行时间
	MaxMem            int                `gorm:"column:max_mem" json:"max_mem"`         // 最大的运行内存
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"`
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) (tx *gorm.DB) {
	tx = DB.Model(new(ProblemBasic)).Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%").Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc ON pc.problem_id = problem_basic.id").Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ?)", categoryIdentity)
	}
	return
}
