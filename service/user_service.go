package service

import (
	"gvf_server/global"
	"gvf_server/models"
)

// GetUserInfo 根据userID查询用户
func GetUserInfo(userID interface{}) (user models.UserModel, err error) {
	err = global.DB.Find(&user, "id = ?", userID).Error
	return
}
