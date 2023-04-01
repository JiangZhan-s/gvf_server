package api

import (
	"gvf_server/api/file_api"
	"gvf_server/api/settings_api"
	"gvf_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	UserApi     user_api.UserApi
	FileApi     file_api.FileApi
}

var ApiGroupApp = new(ApiGroup)