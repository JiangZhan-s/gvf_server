package api

import (
	"gvf_server/api/file_api"
	"gvf_server/api/file_folder_api"
	"gvf_server/api/settings_api"
	"gvf_server/api/share_api"
	"gvf_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi   settings_api.SettingsApi
	UserApi       user_api.UserApi
	FileApi       file_api.FileApi
	FileFolderApi file_folder_api.FileFolderApi
	ShareAPI      share_api.ShareApi
}

var ApiGroupApp = new(ApiGroup)
