package service

import (
	"gvf_server/global"
	"gvf_server/models"
)

// CreateFolderRoot 创建用户根目录
// 根目录名称为用户的用户名，父目录ID为0
func CreateFolderRoot(fileStoreId int, userName string) (*models.FileFolderModel, error) {
	fileFolder := models.FileFolderModel{
		FileFolderName: userName,
		FileStoreID:    fileStoreId,
		ParentFolderID: 0,
	}
	if err := global.DB.Create(&fileFolder).Error; err != nil {
		return nil, err
	}
	return &fileFolder, nil
}
