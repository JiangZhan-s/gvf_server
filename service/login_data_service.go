package service

import (
	"fmt"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/service/common"
)

func AddLoginData(userID int, ip string, nickName string, token string, addr string) string {
	myLogin := models.LoginDataModel{
		UserID:   userID,
		IP:       ip,
		Token:    token,
		Device:   "web",
		Addr:     addr,
		NickName: nickName,
	}
	if err := global.DB.Create(&myLogin).Error; err != nil {
		// 如果保存出错，输出错误信息并退出函数
		fmt.Println("failed to create loginData:", err)
		return ""
	}
	return fmt.Sprintf("%d", myLogin.ID)
}

func GetLoginDataAll(cr models.PageInfo) (loginData []models.LoginDataModel, count int64, err error) {
	searchCond := ""
	var searchValues []interface{}
	loginData, count, err = common.ComList(models.LoginDataModel{}, common.Option{PageInfo: cr}, searchCond, searchValues...)
	return loginData, count, err
}

func GetLoginCount() ([]models.LoginCountResult, error) {
	var loginCounts []models.LoginCountResult

	// 查询数据库中的数据
	err := global.DB.Table("login_data_models").
		Select("DATE(created_at) as date, COUNT(*) as count").
		Group("DATE(created_at)").
		Scan(&loginCounts).Error
	if err != nil {
		return nil, err
	}

	return loginCounts, nil
}
