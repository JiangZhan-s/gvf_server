package file_folder_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
	"os"
	"strconv"
)

func (FileFolderApi) FolderDeleteView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fileId, _ := strconv.Atoi(c.GetHeader("folder_id"))
	folder, err := service.GetFolderById(c.GetHeader("folder_id"))
	if err != nil {
		res.FailWithMessage("删除文件夹失败", c)
		global.Log.Error(err)
		return
	}
	err = service.DeleteFolderById(fileId)
	if err != nil {
		res.FailWithMessage("文件夹下有内容，无法删除", c)
		return
	}
	folderPath := service.GetCurrentFolderPath(folder)
	os.RemoveAll(global.Path + "/" + folderPath)

	res.OkWithMessage("删除成功", c)
}
