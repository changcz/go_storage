package main

import (
	"api-gateway/config"
	"api-gateway/discovery"
	"api-gateway/internal/service"
	"api-gateway/routes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitConfig()

	//服务发现

	etcdAddress := []string{viper.GetString("etcd.address")}

	etcdRegister := discovery.NewResolver(etcdAddress,logrus.New())

	resolver.Register(etcdRegister)

	go startListen()
	{
		osSignal := make(chan os.Signal,1)
		signal.Notify(osSignal,os.Interrupt,os.Kill,syscall.SIGTERM,syscall.SIGINT,syscall.SIGKILL)
		s := <-osSignal
		fmt.Println("exit",s)
	}
	fmt.Println("gateway listen on :4000")
}

func startListen()  {
	 opts :=  []grpc.DialOption{
		grpc.WithInsecure(),
	}
	userConn,err := grpc.Dial("127.0.0.1:10001",opts...)
	if err != nil {
		panic(err)
	}
	userService := service.NewUserServiceClient(userConn)

	taskConn,err := grpc.Dial("127.0.0.1:10002",opts...)
	if err != nil {
		panic(err)
	}
	taskService := service.NewTaskServiceClient(taskConn)

	ginRouter := routes.NewRouter(userService,taskService)
	server := &http.Server{
		Addr: viper.GetString("server.port"),
		Handler: ginRouter,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1<<20, // 服务最大数
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("绑定失败，端口被占用",err)
	}
}