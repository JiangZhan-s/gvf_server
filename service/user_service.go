package service

import (
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/service/common"
)

// GetUserInfo 根据userID查询用户
func GetUserInfo(userID interface{}) (user models.UserModel, err error) {
	err = global.DB.Find(&user, "id = ?", userID).Error
	return
}

// GetUserNickById 根据userID查询用户昵称
func GetUserNickById(userID int) (nickName string, err error) {
	var user models.UserModel
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		// 错误处理
		return "", err
	}
	return user.NickName, err
}

// GetUserAll 获取用户
func GetUserAll(cr models.PageInfo) (users []models.UserModel, count int64, err error) {
	searchCond := ""
	var searchValues []interface{}
	users, count, err = common.ComList(models.UserModel{}, common.Option{PageInfo: cr}, searchCond, searchValues...)
	return users, count, err
}
