package role_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

// RoleQueryAllView 查询所有角色
func (RoleApi) RoleQueryAllView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}

	var cr models.PageInfo
	err = c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//调用搜索服务
	files, count, err := service.GetRoleAll(cr)
	if err != nil {
		res.FailWithMessage("搜索失败", c)
		return
	}
	res.OKWithList(files, count, c)

}
