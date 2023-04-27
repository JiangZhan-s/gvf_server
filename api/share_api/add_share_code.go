package share_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

func (ShareApi) AddShareCodeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fmt.Println(userID)

	fileId := c.GetHeader("file_id")
	fileModel := service.GetFileInfo(fileId)

	//msg, err := global.ServiceSetup.QueryShareCode(fileId)
	//if err != nil {
	//
	//} else {
	//	res.FailWithMessage("wenjian yijing fenxiang ", c)
	//	return
	//}

	msg, err := global.ServiceSetup.StoreShareCode(fileId, fileModel.FileName, string(user.ID))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	msg, err = global.ServiceSetup.QueryShareCode(fileId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("搜索到的区块连值为：", msg)
	}

	//解析json数据
	var d models.Share
	err = json.Unmarshal([]byte(msg), &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	service.ShareFileUp(fileId)
	res.OkWithData(d.ShareCode, c)
}
