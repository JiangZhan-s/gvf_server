package routers

import (
	"gvf_server/api"
	"gvf_server/middleware"
)

func (router RouterGroup) RoleRouter() {
	RoleApi := api.ApiGroupApp.RoleAPI
	router.POST("add_role", middleware.JwtAuth(), RoleApi.AddRoleView)
	router.GET("query_role_all", middleware.JwtAuth(), RoleApi.RoleQueryAllView)
	router.POST("delete_role_by_id", middleware.JwtAuth(), RoleApi.RoleDeleteByIdView)
	router.GET("query_permission_all", middleware.JwtAuth(), RoleApi.PermissionQueryView)
	router.POST("update_role_permission", middleware.JwtAuth(), RoleApi.UpdateRolePermissions)
	router.GET("query_permissions_by_role_id", middleware.JwtAuth(), RoleApi.GetRolePermissions)

}
