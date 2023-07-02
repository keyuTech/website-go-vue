package routers

import (
	"github.com/gin-gonic/gin"
	v1 "keyu.tech/website/api/v1"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingsApi := v1.ApiGroupApp.SettingsApi

	router.GET("settings", settingsApi.SettingsInfoView)
}
