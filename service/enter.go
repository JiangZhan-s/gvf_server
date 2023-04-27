package service

import (
	"gvf_server/service/user_ser"
)

type ServiceGroup struct {
	USerService user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
