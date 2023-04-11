package file_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils"
	"gvf_server/utils/jwts"
	"io"
	"os"
)

func (FileApi) FileUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	folderID := c.GetHeader("id")
	fmt.Println(folderID)
	//接收上传文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		res.FailWithMessage("文件上传错误", c)
		res.FailWithError(err, folderID, c)
		return
	}
	defer file.Close()
	//判断当前文件夹是否有同名文件
	if service.CurrFileExists(folderID, header.Filename) {
		res.FailWithMessage("当前文件夹已有同名文件存在", c)
		return
	}

	//判断用户的容量是否足够
	if !service.CapacityIsEnough(header.Size, user.FileStoreID) {
		res.FailWithMessage("用户容量不足", c)
		return
	}

	newFile, err := os.Create(global.Path + "/" + header.Filename)
	if err != nil {
		res.FailWithMessage("文件创建失败", c)
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		res.FailWithMessage("文件拷贝错误", c)
		return
	}

	//将光标移至开头
	_, err = newFile.Seek(0, 0)
	if err != nil {
		res.FailWithMessage("文件光标移动错误", c)
		return
	}
	hashData := utils.GetSHA256HashCode(newFile)
	fmt.Println(hashData)
	//新建文件信息
	fileID := service.CreateFile(header.Filename, fileSize, folderID, user.FileStoreID, int(user.ID))
	fmt.Println(fileID)
	//上传成功减去相应剩余容量
	service.SubtractSize(fileSize, user.FileStoreID)
	msg, err := global.ServiceSetup.SetInfo(fileID, hashData)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	msg, err = global.ServiceSetup.GetInfo(fileID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	res.OkWithMessage(fmt.Sprintf("文件%s上传成功", header.Filename), c)
}
