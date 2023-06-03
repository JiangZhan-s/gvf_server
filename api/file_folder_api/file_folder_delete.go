package file_folder_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
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
	fileId, err := strconv.Atoi(c.GetHeader("folder_id"))
	err = service.DeleteFolderById(fileId)
	if err != nil {
		res.FailWithMessage("文件夹下有内容，无法删除", c)
		return
	}
	res.OkWithMessage("删除成功", c)
}
