package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
}

// GetMysqlFromConsul 获取mysql 配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}

	// 处理错误返回
	config.Get(path...).Scan(mysqlConfig)

	return mysqlConfig
}
