package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobackend/conf"
	"gobackend/db"
	"gobackend/middleware"
	"gobackend/router"
	"log"
)

func main() {
	// 初始化配置文件
	err := conf.LoadConf(conf.GetConfPath())
	if err != nil {
		log.Fatalln("load config error", err)
	}
	// 取配置
	config := conf.GetConf()

	//初始化Mysql数据库
	_ = db.MysqlDial(&config.Mysql)
	defer db.Eloquent.Close()
	fmt.Println("MySQL connection is successful")

	// 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色。
	gin.DisableConsoleColor()
	//创建空白中间件
	app := gin.New()
	app.Use(middleware.Cors())
	router.Init(app)
	// Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	app.Use(gin.Recovery())

	_ = app.Run(config.Addr) // 在 0.0.0.0:8080 上监听并服务
}
