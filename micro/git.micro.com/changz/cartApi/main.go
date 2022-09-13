package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2"
	"git.micro.com/changz/cartApi/handler"
	"git.micro.com/changz/cartApi/client"

	cartApi "git.micro.com/changz/cartApi/proto/cartApi"
)

func main() {

	//注册中心
	//go get github.com/micro/go-plugins/registry/consul/v2
	//go get github.com/micro/go-micro/v2/registry

	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪


	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the CartApi service client
		micro.WrapHandler(client.CartApiWrapper(service)),
	)

	// Register Handler
	cartApi.RegisterCartApiHandler(service.Server(), new(handler.CartApi))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
