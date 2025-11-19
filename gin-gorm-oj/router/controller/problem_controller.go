package controller

import (
	"gin-gorm-oj/service"

	"github.com/gin-gonic/gin"
)

type ProblemController struct {
}

func (p *ProblemController) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/problem-list", service.GetProblemList)
}
