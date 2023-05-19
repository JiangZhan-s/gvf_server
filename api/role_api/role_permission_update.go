package role_api

import (
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"net/http"
)

func (RoleApi) UpdateRolePermissions(c *gin.Context) {
	var requestData struct {
		RoleID        uint   `json:"roleId"`
		PermissionIDs []uint `json:"permissionIds"`
	}
	// 解析前端发送的请求数据
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	//删除角色所有映射
	err := service.DeleteRolePermissionByRoleId(requestData.RoleID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除角色权限映射失败", c)
	}

	//添加新的映射
	err = service.AddRolePermissionBatch(requestData.RoleID, requestData.PermissionIDs)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("新建角色权限映射失败", c)
	}
	res.OkWithMessage("更新角色权限映射成功！", c)
}
