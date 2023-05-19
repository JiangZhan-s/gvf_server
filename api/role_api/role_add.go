package role_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
	"net/http"
)

func (RoleApi) AddRoleView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	var role models.RoleModel
	if err = c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = service.AddRole(role); err != nil {
		res.FailWithMessage("创建角色失败", c)
	}
	res.OkWithMessage("角色新建成功", c)

}
