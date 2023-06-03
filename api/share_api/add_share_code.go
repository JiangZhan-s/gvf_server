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
	"strconv"
	"time"
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

	fileId := c.GetHeader("file_id")
	fileModel := service.GetFileInfo(fileId)

	var msg string
	var trans string
	maxRetry := 5 // 设置最大重试次数
	for i := 0; i < maxRetry; i++ {
		trans, msg, err = global.ServiceSetup.StoreShareCode(fileId, fileModel.FileName, string(user.ID))
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(trans, msg)
			break // 成功获取到结果，跳出循环
		}
		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

	//解析json数据
	var d models.Share
	err = json.Unmarshal([]byte(msg), &d)
	if err != nil {
		panic(err)
	}
	//分享标志置为1
	service.ShareFileUp(fileId)
	//更新分享表
	hash := service.CreateShare(fileId, strconv.Itoa(int(userID)))
	data := "分享查询码是" + hash + "；分享提取码是" + d.ShareCode
	res.OkWithData(data, c)
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.LogAction(userID, "创建分享", hash)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
	}
}
