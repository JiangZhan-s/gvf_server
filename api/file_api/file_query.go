package file_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

func (FileApi) FileQueryAllView(c *gin.Context) {
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

	var cr models.PageInfo
	err = c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//调用搜索服务
	files, count, err := service.GetUserFileAll(user.FileStoreID, cr, parentFolderId)
	if err != nil {
		res.FailWithMessage("搜索失败", c)
		return
	}
	res.OKWithList(files, count, c)

}

func (FileApi) FIleDetailUseView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	detailUse := service.GetFileDetailUse(user.FileStoreID)
	res.OkWithData(detailUse, c)
}

func (FileApi) FileQueryWithFolderAllView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	parentFolderId := c.GetHeader("parent_folder_id")

	fileWithFolder := service.GetFileWithFolder(parentFolderId)
	res.OkWithData(fileWithFolder, c)
}
