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

func (FileFolderApi) CurrentFolderPathView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	folderId := c.GetHeader("folder_id")
	folder, err := service.GetFolderById(folderId)

	path := service.GetCurrentFolderPath(folder)

	res.OkWithData(path, c)
}

func (FileFolderApi) ParentFolderIdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	folderId := c.GetHeader("folder_id")
	parentFolderId, err := service.GetParentFolderIDById(folderId)

	if parentFolderId == 0 {
		res.FailWithMessage("已经是根文件夹", c)
		return
	}
	res.OkWithData(parentFolderId, c)
}

func (FileFolderApi) FolderCountView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	count, err := service.GetFolderCount(user.FileStoreID)
	res.OkWithData(count, c)
}
