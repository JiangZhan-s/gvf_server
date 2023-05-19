package models

import (
	"gorm.io/gorm"
	"gvf_server/models/ctype"
)

// UserModel User用户表
type UserModel struct {
	gorm.Model
	NickName       string           `gorm:"size:36" json:"nickname"`             //昵称
	UserName       string           `gorm:"size:36" json:"username"`             //用户名
	Password       string           `gorm:"size:128" json:"password"`            //密码
	Email          string           `gorm:"size:128" json:"email"`               //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                  //手机号
	Addr           string           `gorm:"size:64" json:"addr"`                 //地址
	IP             string           `gorm:"size:20" json:"ip"`                   //ip地址
	Role           ctype.Role       `gorm:"size:4" json:"role"`                  //权限：1管理员 2普通用户 3游客
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"` //注册来源
	FileStoreID    int              `json:"file_store_id"`
	FileStoreModel FileStoreModel   `gorm:"foreignKey:UserID" json:"-"`
}
