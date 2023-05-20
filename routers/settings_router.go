package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) SettingsRouter() {
	SettingsApi := api.ApiGroupApp.SettingsApi
	router.GET("query_action", middleware.JwtAuth(), SettingsApi.ActionInfoView)
	router.GET("query_ledger", middleware.JwtAuth(), SettingsApi.LedgerInfoView)
}
