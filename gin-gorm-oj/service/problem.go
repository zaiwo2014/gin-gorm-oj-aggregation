package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "分类标识"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", define.DefaultPageSize))
	if err != nil {
		return
	}

	var count int64

	categoryIdentity := c.Query("category_identity")

	page = page - 1

	keyword := c.DefaultQuery("keyword", "")

	tx := models.GetProblemList(keyword, categoryIdentity)
	list := make([]*models.ProblemBasic, 0)

	err = tx.Count(&count).Omit("content").Limit(pageSize).Offset(page * pageSize).Find(&list).Error

	if err != nil {
		log.Printf("查询失败,失败原因: %v \n", err)
		return
	}

	c.JSON(200, gin.H{
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		}})
}
