package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) FileRouter() {
	app := api.ApiGroupApp.FileApi
	router.POST("upload", middleware.JwtAuth(), app.FileUploadView)
	router.POST("upload_multi", middleware.JwtAuth(), app.MultiFileUploadView)
	router.GET("query_all", middleware.JwtAuth(), app.FileQueryAllView)
	router.GET("download", middleware.JwtAuth(), app.FileDownloadByIdView)
	router.GET("query_detail_use", middleware.JwtAuth(), app.FIleDetailUseView)
	router.GET("query_file_with_folder", middleware.JwtAuth(), app.FileQueryWithFolderAllView)
	router.POST("upload_folder", middleware.JwtAuth(), app.FolderUploadView)
	router.POST("file_delete", middleware.JwtAuth(), app.FileDeleteView)
}
