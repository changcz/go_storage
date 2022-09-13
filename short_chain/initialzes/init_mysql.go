package initialzes

import (
	"fmt"
	"goproject/short_chain/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{},
	)

	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.SysConfig.MysqlConfig.User,
		global.SysConfig.MysqlConfig.Password,
		global.SysConfig.MysqlConfig.Host,
		global.SysConfig.MysqlConfig.Port,
		global.SysConfig.MysqlConfig.DbName,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		//DSN:                       "heitao:O8ZlWq9kh6pSeEtu@tcp(10.55.10.88:6033)/db_heitao_act_yhdzz?charset=utf8mb4&parseTime=True&loc=Local",
		DSN:                       mysqlUrl,
		DefaultStringSize:         256,  // string类型字段默认长度
		DisableDatetimePrecision:  true, // 禁止datetime精度，mysql 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引，就要把索引先删除再重建mysql 5.7 不支持
		DontSupportRenameColumn:   true, // 用change 重命名索引，mysql 8 之前的数据不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, // 命名策略 不加s
	})
	if err != nil {
		log.Fatalf("Open err:%v", err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 空闲连接数
	sqlDB.SetMaxOpenConns(100)                 // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 超时

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Ping err:%v", err)
		return
	}

	global.Db = db
	fmt.Println("DB success")
}
