package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) LoginDataRouter() {
	LoginDataApi := api.ApiGroupApp.LoginDataApi
	router.GET("query_login_data_all", middleware.JwtAuth(), LoginDataApi.LoginDataQueryAllView)
	router.GET("query_login_count", LoginDataApi.LoginCountQueryView)
}
