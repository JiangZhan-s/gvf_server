package common

import (
	"gorm.io/gorm"
	"gvf_server/global"
	"gvf_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option, searchCond string, searchValues ...interface{}) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	count = DB.Model(model).Select("id").Where(searchCond, searchValues...).Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Model(model).Where(searchCond, searchValues...).Limit(option.Limit).Offset(offset).Find(&list).Error
	return list, count, err
}
