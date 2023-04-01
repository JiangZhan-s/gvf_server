package models

import "gorm.io/gorm"

// FileFolderModel 文件夹表
type FileFolderModel struct {
	gorm.Model
	FileFolderName string         `json:"file_folder_name"` //文件夹名
	ParentFolderID int            `json:"parent_folder_id"` //父文件夹ID
	FileStoreID    int            `json:"file_store_id"`    //文件所属仓库ID
	Time           string         `json:"time"`             //时间
	FileStoreModel FileStoreModel `gorm:"foreignKey:FileStoreID" json:"file_store"`
	FilesModel     []FileModel    `gorm:"foreignKey:ParentFolderID" json:"file_model"`
}
