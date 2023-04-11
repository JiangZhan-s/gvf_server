package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) FileRouter() {
	app := api.ApiGroupApp.FileApi
	router.POST("upload", middleware.JwtAuth(), app.FileUploadView)
	router.GET("query_all", middleware.JwtAuth(), app.FileQueryAllView)
	router.GET("download", middleware.JwtAuth(), app.FileDownloadByIdView)
}
