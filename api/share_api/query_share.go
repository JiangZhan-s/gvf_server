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
	"time"
)

// ShareQueryByIdView 根据文件id在区块连中搜索提取码
func (ShareApi) ShareQueryByIdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fmt.Println(user.ID)

	fileId := c.GetHeader("file_id")

	var msg string
	maxRetry := 5 // 设置最大重试次数
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.QueryShareCode(fileId)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
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
	fmt.Println(d)
	service.ShareFileUp(fileId)
	res.OkWithData(d.ShareCode, c)

}

// CodeQueryByIdView 根据文件id搜索查询码
func (ShareApi) CodeQueryByIdView(c *gin.Context) {
	fileId := c.GetHeader("file_id")
	var code models.ShareModel
	global.DB.Where("file_id = ?", fileId).First(&code)
	res.OkWithData(code.Hash, c)
}

// ShareInfoQueryByHash 根据查询码查询文件信息
func (ShareApi) ShareInfoQueryByHash(c *gin.Context) {
	shareHash := c.GetHeader("share_hash")
	share := service.GetShareByHash(shareHash)
	if share.ID == 0 {
		res.FailWithMessage("无效的查询码", c)
	}
	res.OkWithData(share, c)

}

// FileInfoQueryByCode 验证提取玛并返回文件信息或错误信息
func (ShareApi) FileInfoQueryByCode(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	hashCode := c.GetHeader("hash_code")
	fileId := c.GetHeader("file_id")

	//获取区块链中的hashcode
	var msg string
	maxRetry := 5 // 设置最大重试次数
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.QueryShareCode(fileId)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

	//解析json数据
	var d models.Share
	var success string
	err = json.Unmarshal([]byte(msg), &d)
	if err != nil {
		panic(err)
	}
	if d.ShareCode == hashCode {
		fileInfo := service.GetFileInfo(fileId)
		res.OkWithData(fileInfo, c)
		success = "成功"
		return
	} else {
		res.FailWithMessage("错误的提取码", c)
		success = "失败"
	}
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.LogAction(user.ID, "提取文件"+success, "文件id:"+fileId)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
	}
}

func (ShareApi) ShareCountView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	count, err := service.GetShareCount(userID)
	res.OkWithData(count, c)
}
