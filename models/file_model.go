package models

import "gorm.io/gorm"

// FileModel 文件表
type FileModel struct {
	gorm.Model
	UserID         int            `json:"user_id"`                                  //用户id
	UserModel      UserModel      `gorm:"foreignKey:UserID" json:"user_model"`      //添加外键关联到UserModel
	FileName       string         `json:"file_name"`                                //文件名
	FileStoreID    int            `json:"file_store_id"`                            //文件仓库id
	FilePath       string         `json:"file_path"`                                //文件存储路径
	DownloadNum    int            `json:"download_num"`                             //下载次数
	UploadTime     string         `json:"upload_time"`                              //上传时间
	ParentFolderID int            `json:"parent_folder_id"`                         //父文件夹id
	Size           int64          `json:"size"`                                     //文件大小
	SizeStr        string         `json:"size_str"`                                 //文件大小单位
	Type           int            `json:"type'"`                                    //文件类型
	Postfix        string         `json:"postfix"`                                  //文件后缀
	FileStoreModel FileStoreModel `gorm:"foreignKey:FileStoreID" json:"file_store"` // 添加外键关联到FileStoreModel
}

type Data struct {
	DataHash string `json:"DataHash"`
}

type Share struct {
	OwnerId   string `json:"OwnerId"`
	ShareCode string `json:"ShareCode"`
}
