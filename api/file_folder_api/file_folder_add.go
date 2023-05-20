package file_folder_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

// FolderAddView 创建文件夹
func (FileFolderApi) FolderAddView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	parentFolderId := c.GetHeader("parent_folder_id")
	folderName := c.GetHeader("folder_name")
	service.CreateFolder(folderName, parentFolderId, user.FileStoreID)
	res.OkWithMessage("创建成功", c)
	maxRetry := 5
	var msg string
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.LogAction(userID, "新建文件夹", folderName)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
	}
}
