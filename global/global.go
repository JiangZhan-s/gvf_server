package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvf_server/config"
)

var (
	Config        *config.Config
	DB            *gorm.DB
	Log           *logrus.Logger
	Path          string
	ServiceSetup  config.ServiceSetup
	ChannelClient *channel.Client
	MysqlLog      logger.Interface
	Redis         *redis.Client
)
