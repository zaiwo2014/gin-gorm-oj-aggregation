package controller

import (
	"gin-gorm-oj/service"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (pingController *PingController) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/ping", service.Ping)
}
