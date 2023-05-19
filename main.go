package main

import (
	"gvf_server/core"
	"gvf_server/flag"
	"gvf_server/global"
	"gvf_server/routers"
)

func main() {
	var err error
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
	//初始化fabric
	//global.ChannelClient, global.ServiceSetup, err = core.InitFabric()
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addrr()
	global.Log.Infof("gvf_server运行在:%s", addr)
	err = router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
