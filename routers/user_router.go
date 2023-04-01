package routers

import (
	"github.com/gin-gonic/gin"
	"gvf_server/api"
)

func UserRouter(router *gin.Engine) {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
}
