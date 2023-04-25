package service

import (
	"gvf_server/global"
	"gvf_server/models"
	"os"
	"path/filepath"
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
	folderPath := filepath.Join(global.Path, userName)
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		panic(err)
	}
	return &fileFolder, nil
}

// FindFolderRoot 查询用户根目录
// 根据用户所对应的仓库id查询用户的根目录(根目录的父id为0)
func FindFolderRoot(fileStoreId int) (int, error) {
	result := -1
	if err := global.DB.Model(&models.FileFolderModel{}).Select("id").Where("file_store_id=? and parent_folder_id=?", fileStoreId, 0).Scan(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}
