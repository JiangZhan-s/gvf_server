package api

import (
	"gvf_server/api/file_api"
	"gvf_server/api/file_folder_api"
	"gvf_server/api/file_store_api"
	"gvf_server/api/login_data_api"
	"gvf_server/api/role_api"
	"gvf_server/api/settings_api"
	"gvf_server/api/share_api"
	"gvf_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi   settings_api.SettingsApi
	UserApi       user_api.UserApi
	FileApi       file_api.FileApi
	FileFolderApi file_folder_api.FileFolderApi
	FileStoreApi  file_store_api.FileStoreApi
	ShareApi      share_api.ShareApi
	RoleApi       role_api.RoleApi
	LoginDataApi  login_data_api.LoginDataApi
}

var ApiGroupApp = new(ApiGroup)
