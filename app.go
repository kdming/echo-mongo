package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mongo "web_model/src/common/mongo"
	router "web_model/src/common/router"
)

func main() {
	// 建立数据库连接
	isConnect := mongo.Connect()
	if isConnect == true {
		e := echo.New()
		// 开启中间件
		e.Use(middleware.CORS())
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		// 加载路由
		router.InitRouter(e)
		e.Logger.Fatal(e.Start(":46200"))
	}
}
