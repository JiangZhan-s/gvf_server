package routers

import (
	"github.com/gin-gonic/gin"
	"gvf_server/api"
	"gvf_server/middleware"
)

func FileRouter(router *gin.Engine) {
	app := api.ApiGroupApp.FileApi
	router.POST("upload", middleware.JwtAuth(), app.FileUploadView)
}
