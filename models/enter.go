package models

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type FolderInfo struct {
	FolderName     string `json:"folderName"`
	ParentFolderID int    `json:"parentFolderId"`
}
