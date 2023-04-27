package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) ShareRouter() {
	app := api.ApiGroupApp.ShareAPI
	router.POST("share_generate", middleware.JwtAuth(), app.AddShareCodeView)
	router.GET("query_share_all", middleware.JwtAuth(), app.FileQueryAllView)
	router.GET("query_share_by_id", middleware.JwtAuth(), app.FileQueryByIdView)
}
