package routers

import (
	"github.com/gin-gonic/gin"
	"gvf_server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//系统配置api
	SettingsRouter(router)
	UserRouter(router)
	FileRouter(router)
	return router
}
