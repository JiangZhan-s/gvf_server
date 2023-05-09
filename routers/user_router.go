package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.GET("get_nickname_by_id", app.UserNickQueryByIdView)
}
