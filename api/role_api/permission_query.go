package role_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

// PermissionQueryView 查询所有权限
func (RoleApi) PermissionQueryView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}

	var permissions []models.PermissionModel
	//调用搜索服务
	permissions, err = service.GetPermissionAll()
	if err != nil {
		res.FailWithMessage("搜索失败", c)
		return
	}
	res.OkWithData(permissions, c)

}

// GetRolePermissions 处理前端发送的请求来查询角色的所有映射
func (RoleApi) GetRolePermissions(c *gin.Context) {
	roleID := c.Query("roleId")
	fmt.Println(roleID)

	var rolePermissions []models.RolePermissionModel
	if err := global.DB.Select("permission_id").Where("role_id = ?", roleID).Find(&rolePermissions).Error; err != nil {
		res.FailWithMessage("查询映射失败", c)
		return
	}

	permissionIDs := make([]uint, len(rolePermissions))
	for i, rp := range rolePermissions {
		permissionIDs[i] = rp.PermissionID
	}

	res.OkWithData(permissionIDs, c)
}
