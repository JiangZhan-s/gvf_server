package global

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvf_server/config"
)

var (
	Config        *config.Config
	DB            *gorm.DB
	Log           *logrus.Logger
	Path          string
	ServiceSetup  config.ServiceSetup
	ChannelClient *channel.Client
)
