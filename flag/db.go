package flag

import (
	"gvf_server/global"
	"gvf_server/models"
)

func MakeMigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserModel{})
	//生成表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.LoginDataModel{},
			&models.FileModel{},
			&models.FileFolderModel{},
			&models.FileStoreModel{},
			&models.ShareModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功")
}
