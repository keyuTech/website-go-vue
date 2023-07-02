package routers

import (
	"github.com/gin-gonic/gin"
	// swaggerFile "github.com/swaggo/files"
	// gs "github.com/swaggo/gin-swagger"
	"keyu.tech/website/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// Router.GET("/swagger/*any", gs.WrapHandler(swaggerFile.Handler))
	PublicGroup := Router.Group("v1")

	SettingsRouter(PublicGroup)
	return Router
}
