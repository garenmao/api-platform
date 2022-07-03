package main

import (
	"api-platform/initialize"
)

func main() {
	// 初始化配置
	initialize.Configs()
	// 初始化logger
	initialize.Logger()
	// 初始化redis
	initialize.Redis()
	// 初始化mysql
	initialize.Mysql()
	//global.DB.AutoMigrate(domain.User{})
	engine := initialize.Routers()
	engine.Run(":8080")
}
