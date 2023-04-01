package service

import (
	"gvf_server/global"
	"gvf_server/models"
)

// CreateFileStore 创建文件仓库
func CreateFileStore(userId int) (*models.FileStoreModel, error) {
	fileStore := models.FileStoreModel{
		UserID:      userId,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	if err := global.DB.Create(&fileStore).Error; err != nil {
		return nil, err
	}
	return &fileStore, nil
}

// GetUserFileStore 根据用户id获取仓库信息
func GetUserFileStore(userId int) (*models.FileStoreModel, error) {
	var fileStore models.FileStoreModel
	if err := global.DB.Find(&fileStore, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &fileStore, nil
}

// CapacityIsEnough 判断用户容量是否足够
func CapacityIsEnough(fileSize int64, userId int) bool {
	fileStore, err := GetUserFileStore(userId)
	if err != nil {
		return false
	}

	remainingSize := fileStore.MaxSize - fileStore.CurrentSize - fileSize
	return remainingSize >= 0
}
