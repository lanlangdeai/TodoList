package main

import (
	"to-do-list/conf"
	"to-do-list/routes"
)

// @title ToDoList API
// @version 0.0.1
// @description This is todo_list project
// @content.name lixing02
// @content.url http://localhost:3005/swagger/index.html
// @content.email lixing02@bianfeng.com

// @host localhost:3005
// @BasePath /api/v1
func main() {
	//从配置文件读入配置
	conf.Init()
	//转载路由 swag init -g common.go
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
