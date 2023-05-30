package main

import (
	"fmt"
	"gvf_server/config"
	"gvf_server/core"
	"gvf_server/flag"
	"gvf_server/global"
	"gvf_server/routers"
	"gvf_server/sdkInit"
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
	//连接redis
	global.Redis = core.ConnectRedis()
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

	//err = sdkInit.CreateChannel(sdk, initInfo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//global.ChannelClient, err = sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(global.ChannelClient)

	global.ChannelClient, err = sdkInit.GetChannelClient(sdk, initInfo)
	if err != nil {
		fmt.Printf("获取通道客户端失败: %v", err)
		return
	}

	// 在这里可以使用 channelClient 对象调用链码进行查询或执行事务
	// ...

	global.ServiceSetup = config.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      global.ChannelClient,
	}

	router := routers.InitRouter()
	addr := global.Config.System.Addrr()
	global.Log.Infof("gvf_server运行在:%s", addr)
	err = router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
