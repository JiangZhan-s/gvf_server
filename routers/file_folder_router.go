package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) FileFolderRouter() {
	app := api.ApiGroupApp.FileFolderApi
	router.GET("folder_root_find", middleware.JwtAuth(), app.FolderRootFindView)
	router.POST("add_file_folder", middleware.JwtAuth(), app.FolderAddView)
	router.GET("query_parent_folder_id", middleware.JwtAuth(), app.ParentFolderIdView)
	router.POST("folder_delete", middleware.JwtAuth(), app.FolderDeleteView)
	router.GET("query_folder_count", middleware.JwtAuth(), app.FolderCountView)
}
