package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd() // 获取工作目录
	viper.SetConfigName("config") //  配置文件名
	viper.SetConfigType("yml") // 配置文件后缀
	viper.AddConfigPath(workDir + "/config") 
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}