package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) FileRouter() {
	app := api.ApiGroupApp.FileApi
	router.POST("upload", middleware.JwtAuth(), app.FileUploadView)
}
