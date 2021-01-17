package main

import (
	"gin-demo/global"
	"gin-demo/server"
	"gin-demo/router"
	db "gin-demo/database"
)

// @title 文档
// @version 0.1
// @description 统计 api
// @termsOfService http://swagger.io/terms/

// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 
// @BasePath /
func main() {
	
	global.VP = server.Viper() // 初始化Viper
	global.LOG = server.Zap()           // 初始化zap日志库
	global.DB = db.Gorm()     // gorm连接数据库
	db.MysqlTables(global.DB) // 初始化表

	// 程序结束前关闭数据库链接
	db, _ := global.DB.DB()
	defer db.Close()
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// server.RunWindowsServer()
}
