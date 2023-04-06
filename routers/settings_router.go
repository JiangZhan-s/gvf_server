package routers

import (
	"gvf_server/api"
)

func (router RouterGroup) SettingsRouter() {
	SettingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", SettingsApi.SettingsInfoView)
}
