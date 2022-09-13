package main

import (
	"git.micro.service.category/common"
	"git.micro.service.category/domain/repository"
	service2 "git.micro.service.category/domain/service"
	"git.micro.service.category/handler"
	category "git.micro.service.category/proto/category"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		log.Error(err)
	}

	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New service
	srv := micro.NewService(
		micro.Name("git.micro.com.changz.category"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul 作为注册中心
		micro.Registry(consulRegistry),
	)

	// 获取mysql配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	// root:123456@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local
	//dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@tcp(" + mysqlInfo.Host + ":" + strconv.FormatInt(mysqlInfo.Port, 10) + ")/" + mysqlInfo.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@/" + mysqlInfo.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Error(err)
	}

	defer db.Close()

	// 禁止数据库名复数
	db.SingularTable(true)

	srv.Init()

	rp := repository.NewCategoryRepository(db)
	rp.InitTable() // 创建表  只执行一次
	categoryDataService := service2.NewCategoryRepository(rp)

	// Register handler
	err = category.RegisterCategoryHandler(srv.Server(), &handler.Category{
		CategoryDataService: categoryDataService,
	})
	if err != nil {
		logger.Error(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
