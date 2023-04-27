package user_api

import (
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")
	err := service.ServiceApp.USerService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
