package models

import (
	"gorm.io/gorm"
)

// LoginDataModel  统计用户登陆数据 id,用户id，用户昵称，用户token，登陆设备，登陆时间
type LoginDataModel struct {
	gorm.Model
	UserID    int       `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"`
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"`
	Addr      string    `gorm:"size:64" json:"addr"`
}
