package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) ShareRouter() {
	app := api.ApiGroupApp.ShareApi
	router.POST("share_generate", middleware.JwtAuth(), app.AddShareCodeView)
	router.GET("query_share_all", middleware.JwtAuth(), app.FileQueryAllView)
	router.GET("query_share_by_id", middleware.JwtAuth(), app.ShareQueryByIdView)
	router.GET("get_share_file_id_by_hash", app.ShareInfoQueryByHash)
	router.GET("get_share_file_info_by_code", middleware.JwtAuth(), app.FileInfoQueryByCode)
	router.GET("get_code_by_id", app.CodeQueryByIdView)
}
