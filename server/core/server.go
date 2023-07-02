package core

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"keyu.tech/website/config"
	"keyu.tech/website/global"
)

// 读取yaml配置文件
func InitServer() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
