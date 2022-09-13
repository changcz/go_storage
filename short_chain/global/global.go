package global

import (
	"goproject/short_chain/config"
	"gorm.io/gorm"
)

var (
	SysConfig *config.SysConfig
	Db        *gorm.DB
)
