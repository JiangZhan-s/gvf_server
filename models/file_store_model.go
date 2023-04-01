package models

import "gorm.io/gorm"

// FileStoreModel 文件仓库
type FileStoreModel struct {
	gorm.Model
	UserID           int               `json:"user_id"`                                    //所属用户ID
	CurrentSize      int64             `json:"current_size"`                               //已用空间
	MaxSize          int64             `json:"max_size"`                                   //最大空间
	FileFoldersModel []FileFolderModel `gorm:"foreignKey:FileStoreID" json:"file_folders"` // 添加外键关联到FileFolderModel，并添加json字段
	FilesModel       []FileModel       `gorm:"foreignKey:FileStoreID" json:"files"`        // 添加外键关联到FileModel，并添加json字段
}
