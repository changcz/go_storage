package initialzes

import (
	"gopkg.in/yaml.v2"
	"goproject/short_chain/config"
	"goproject/short_chain/global"
	"io/ioutil"
	"log"
)

func InitConfig() {
	configs := new(config.SysConfig)
	yamlFile, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, configs)
	if err != nil {
		log.Fatalf("Unmarshal:%v", err)
	}
	global.SysConfig = configs
}
