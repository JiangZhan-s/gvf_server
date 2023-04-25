package routers

import (
	"github.com/gin-gonic/gin"
	"gvf_server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApp.SettingsRouter()
	routerGroupApp.UserRouter()
	routerGroupApp.FileRouter()
	routerGroupApp.FileFolderRouter()
	routerGroupApp.ShareRouter()
	return router
}
