package user_api

import (
	"github.com/gin-gonic/gin"
	"gvf_server/models/res"
	"gvf_server/service"
	"strconv"
)

// UserNickQueryByIdView 通过ID查询用户昵称
func (UserApi) UserNickQueryByIdView(c *gin.Context) {
	userId, _ := strconv.Atoi(c.GetHeader("user_id"))
	nickName, _ := service.GetUserNickById(userId)
	res.OkWithData(nickName, c)
}
