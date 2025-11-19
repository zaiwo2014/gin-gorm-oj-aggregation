package router

import (
	_ "gin-gorm-oj/docs"
	"gin-gorm-oj/router/controller"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller interface {
	RegisterRoutes(*gin.Engine)
}

var controllers []Controller

func RegisterController(c Controller) {
	controllers = append(controllers, c)
}

func init() {
	RegisterController(&controller.PingController{})
	RegisterController(&controller.ProblemController{})
}

func Router() *gin.Engine {
	r := gin.Default()

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//注册 controller
	for _, ctrl := range controllers {
		ctrl.RegisterRoutes(r)
	}

	return r

}
