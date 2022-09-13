package repository

import (
	"fmt"
	"os"
)

// 迁移

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&Task{},
		)
	if err != nil {
		fmt.Println("register table fail")
		//util.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}

	fmt.Println("register table success")
	//util.LogrusObj.Infoln("register table success")
}
