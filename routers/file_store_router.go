package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) FileStoreRouter() {
	app := api.ApiGroupApp.FileStoreApi
	router.GET("query_store_by_user_id", middleware.JwtAuth(), app.StoreQueryByUserIdView)
}
