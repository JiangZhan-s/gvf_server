package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvf_server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Path   string
)
