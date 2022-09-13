**gin+grpc+gorm+etcd+mysql 的备忘录功能**


# **api网关**
api-gateway/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑
│   ├── handler           // 视图层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
├── middleware            // 中间件
├── pkg                   // 各种包
│   ├── e                 // 统一错误状态码
│   ├── res               // 统一response接口返回
│   └── util              // 各种工具、JWT、Logger等等..
├── routes                // http路由模块

# **user && task 用户与任务模块**
user/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   ├── cache             // 缓存模块
│   ├── repository        // 持久层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
└── pkg                   // 各种包
├── e                 // 统一错误状态码
├── res               // 统一response接口返回
└── util              // 各种工具、JWT、Logger等等.


各模块下的config/config.yml文件模板

server:
# 模块
domain: user
# 模块名称
version: 1.0
# 模块版本
grpcAddress: "127.0.0.1:10001"
# grpc地址

datasource:
# mysql数据源
driverName: mysql
host: 127.0.0.1
port: 3306
database: grpc_todolist
# 数据库名
username: root
password: root
charset: utf8mb4

etcd:
# etcd 配置
address: 127.0.0.1:2379

redis:
# redis 配置
address: 127.0.0.1:6379
password:


# 微服务开发流程  例如 user 模块

user 模块 主要实现 用户注册、登录

## cmd/main.go为项目的入口文件
## config/config.go  为项目的配置文件
    config/config.yml 具体写配置信息
    通过 viper包读取 yml 文件中的配置信息
    在 main.go 中 调用 InitConfig方法 加载配置信息
## internal/service/pb 中定义 protobuf 文件
    option go_package= "/internal/service;service";
    表示  在 /internal/service 下生成 包名为service 的 proto 文件 
## internal/repository 持久层中定义 mysql链接、数据库模型
    db_int.go 定于数据库 连接配置  在main.go 中  repository.InitDB() 加载
    migration.go 用与迁移 映射 数据库库表
    user.go 定义表结构、操作表
## handler 定义结构体  实现 proto文件中的接口
    比如  NewUserService 实例化 UserService 结构体方法 在 main中 绑定service 服务
	service.RegisterUserServiceServer(server, handler.NewUserService())
    
## discovery etcd服务注册、keep-alive、获取服务信息
    

