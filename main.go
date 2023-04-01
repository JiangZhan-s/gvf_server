package main

import (
	"fmt"
	"gvf_server/core"
	"gvf_server/flag"
	"gvf_server/global"
	"gvf_server/routers"
	"gvf_server/sdkInit"
	"gvf_server/service"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	SimpleCC    = "myapp"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	global.Path = core.InitPath()
	global.Log.Infof("文件根目录路径为:%s", global.Path)
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "mychannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/gvf_project/gvf_server/fixtures/channel-artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.example.com",

		ChaincodeID:     SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "gvf_project/gvf_server/chaincode/",
		UserName:        "Admin",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	serviceSetup := service.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      channelClient,
	}

	msg, err := serviceSetup.SetInfo("1", "123456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	msg, err = serviceSetup.GetInfo("1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addrr()
	global.Log.Infof("gvf_server运行在:%s", addr)
	err = router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
