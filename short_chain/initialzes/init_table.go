package initialzes

import (
	"fmt"
	"goproject/short_chain/global"
	"goproject/short_chain/models/table"
	"log"
	"os"
)

func InitTable() {
	fmt.Println("InitTable")
	db := global.Db
	err := db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&table.ShortChainData{})
	if err != nil {
		log.Printf("ShortChainData err %v ", err)
		os.Exit(0)
	}
	fmt.Println("table success")
}
