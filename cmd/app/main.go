package main

import (
	"luoqiangGin/configs"
	"luoqiangGin/internal/app"
	"luoqiangGin/internal/app/database"
)

//项目启动
func main() {
	configs.Init()
	// 连接数据库
	database.OpenDB()
	// 设置连接池
	database.SetPool()
	app.Init()
}
