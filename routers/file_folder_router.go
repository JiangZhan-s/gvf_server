package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) FileFolderRouter() {
	app := api.ApiGroupApp.FileFolderApi
	router.GET("folder_root_find", middleware.JwtAuth(), app.FolderRootFindView)
	router.POST("add_file_folder", middleware.JwtAuth(), app.FolderAddView)
}
