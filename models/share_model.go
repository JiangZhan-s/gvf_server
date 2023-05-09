package models

import "gorm.io/gorm"

// ShareModel 分享表
type ShareModel struct {
	gorm.Model
	UserID int    `json:"user_id"` //用户id
	FileID int    `json:"file_id"`
	Hash   string `json:"hash"`
}
