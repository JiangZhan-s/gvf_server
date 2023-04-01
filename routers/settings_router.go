package routers

import (
	"github.com/gin-gonic/gin"
	"gvf_server/api"
)

func SettingsRouter(router *gin.Engine) {
	SettingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", SettingsApi.SettingsInfoView)
}
