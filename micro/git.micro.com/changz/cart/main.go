package main

import (
	"git.micro.com/changz/cart/domain/repository"
	service2 "git.micro.com/changz/cart/domain/service"
	"git.micro.com/changz/cart/handler"
	cart "git.micro.com/changz/cart/proto/cart"
	"git.micro.com/changz/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

// QPS 限流数
var QPS = 100

func main() {

	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		log.Error(err)
	}

	//注册中心
	//go get github.com/micro/go-plugins/registry/consul/v2
	// go get github.com/micro/go-micro/v2/registry
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("git.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
	)

	// 数据库链接
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
	rp := repository.NewCartRepository(db)
	//rp.InitTable() // 创建表  只执行一次NewProductRepository

	// New service
	srv := micro.NewService(
		micro.Name("git.micro.service.cart"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("0.0.0.0:8087"),
		// 添加consul 作为注册中心
		micro.Registry(consulRegistry),

		// 绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),

		// 添加熔断限流  go get github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	cartDataService := service2.NewCartRepository(rp)
	// Register handler
	err = cart.RegisterCartHandler(srv.Server(), &handler.Cart{
		CartDataService: cartDataService,
	})
	if err != nil {
		log.Error(err)
	}
	srv.Init()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
