package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) ShareRouter() {
	app := api.ApiGroupApp.ShareAPI
	router.POST("share_generate", middleware.JwtAuth(), app.AddShareCode)
}
