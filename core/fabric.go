package core

//
//const (
//	configFile  = "config.yaml"
//	initialized = false
//	SimpleCC    = "myapp"
//)
//
//func InitFabric() (*channel.Client, config.ServiceSetup, error) {
//
//	initInfo := &sdkInit.InitInfo{
//
//		ChannelID:     "mychannel",
//		ChannelConfig: os.Getenv("GOPATH") + "/src/gvf_project/gvf_server/fixtures/channel-artifacts/channel.tx",
//
//		OrgAdmin:       "Admin",
//		OrgName:        "Org1",
//		OrdererOrgName: "orderer.example.com",
//
//		ChaincodeID:     SimpleCC,
//		ChaincodeGoPath: os.Getenv("GOPATH"),
//		ChaincodePath:   "gvf_project/gvf_server/chaincode/",
//		UserName:        "Admin",
//	}
//	sdk, err := sdkInit.SetupSDK(configFile, initialized)
//	if err != nil {
//		fmt.Printf(err.Error())
//		return nil, config.ServiceSetup{}, err
//	}
//	defer sdk.Close()
//	err = sdkInit.CreateChannel(sdk, initInfo)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil, config.ServiceSetup{}, err
//	}
//	global.ChannelClient, err = sdkInit.InstallAndInstantiateCC(sdk, initInfo)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil, config.ServiceSetup{}, err
//	}
//	fmt.Println(global.ChannelClient)
//	global.ServiceSetup = config.ServiceSetup{
//		ChaincodeID: SimpleCC,
//		Client:      global.ChannelClient,
//	}
//	return global.ChannelClient, global.ServiceSetup, err
//}
