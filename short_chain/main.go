package main

import "goproject/short_chain/initialzes"

func main() {
	initialzes.RunServe()
}

/*
config 配置文件解析
controllers 控制层
global 全局变量
initializes 初始化层
models 模型层 数据库相关的表
router 路由层
services 服务层 处理相关的业务逻辑
config.yaml 配置文件
main.go 主方法入口
*/
