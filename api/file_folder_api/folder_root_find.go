package file_folder_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

func (FileFolderApi) FolderRootFindView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	folderRootId, err := service.FindFolderRoot(user.FileStoreID)
	if err != nil {
		res.FailWithMessage("查找用户文件根目录时出错", c)
		return
	}
	res.OkWithData(folderRootId, c)
}
