package main

import (
	"git.micro.com/changz/pruduct/common"
	"git.micro.com/changz/pruduct/domain/repository"
	"git.micro.com/changz/pruduct/domain/service"
	"git.micro.com/changz/pruduct/handler"
	product "git.micro.com/changz/pruduct/proto/product"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		log.Error(err)
	}

	//注册中心 //go get github.com/micro/go-plugins/registry/consul/v2
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("git.micro.com.changz.product", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 获取mysql配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	// root:123456@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local
	dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@/" + mysqlInfo.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Error(err)
	}

	defer db.Close()

	// 禁止数据库名复数
	db.SingularTable(true)

	rp := repository.NewProductRepository(db)
	//rp.InitTable() // 创建表  只执行一次NewProductRepository

	// New service
	srv := micro.NewService(
		micro.Name("git.micro.com.changz.product"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul 作为注册中心
		micro.Registry(consulRegistry),

		// 绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	productDataService := service.NewProductRepository(rp)
	// Register handler
	err = product.RegisterProductHandler(srv.Server(), &handler.Product{
		ProductDataService: productDataService,
	})
	if err != nil {
		logger.Error(err)
	}

	srv.Init()
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
