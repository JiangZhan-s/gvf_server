package role_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
	"net/http"
)

// RoleDeleteByIdView 根据ID删除角色
func (RoleApi) RoleDeleteByIdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}

	var roleId uint
	if err = c.ShouldBindJSON(&roleId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(roleId)

	//调用删除服务
	roleName, err := service.DeleteRoleById(roleId)
	if err != nil {
		res.FailWithMessage("删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("角色 %s 删除成功", roleName), c)

}
