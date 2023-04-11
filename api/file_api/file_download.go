package file_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils"
	"gvf_server/utils/jwts"
	"os"
	"path"
)

func (FileApi) FileDownloadByIdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fmt.Println(user.IP)

	fileId := c.GetHeader("file_id")
	fileModel := models.FileModel{}
	fileModel = service.GetFileInfo(fileId)
	fmt.Println(fmt.Sprintf(global.Path + "/" + fileModel.FileName + fileModel.Postfix))
	file, err := os.Open(global.Path + "/" + fileModel.FileName + fileModel.Postfix)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer file.Close()
	fmt.Println("文件打开成功")

	hashData := utils.GetSHA256HashCode(file)
	fmt.Println(hashData)
	//对比链上连下数据哈希是否相等
	msg, err := global.ServiceSetup.GetInfo(fileId)
	if err != nil {
		fmt.Println(err)
		return
	}

	//解析json数据
	var d models.Data
	err = json.Unmarshal([]byte(msg), &d)
	if err != nil {
		panic(err)
	}

	if d.DataHash != hashData {
		err = errors.New("链上连下数据协同验证不统一，下载请求被拒绝")
		c.Status(550)
		return
	}
	fmt.Println(msg)
	// 将文件作为附件返回r给客户端进行下载
	filePath := path.Join(global.Path, fileModel.FileName+fileModel.Postfix)
	fileName := fileModel.FileName + fileModel.Postfix
	// 调用函数传输文件
	c.FileAttachment(filePath, fileName)

	// 更新下载次数
	service.DownloadNumAdd(fileId)

}
