package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvf_server/config"
	"gvf_server/global"
	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	fmt.Println(c)
	global.Config = c
}

//go get gopkg.in/yaml.v2
