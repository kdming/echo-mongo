package main

import (
	"web_model/dao"
	"web_model/routers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// 建立数据库连接
	isConnect := dao.Connect()
	if isConnect == true {
		e := echo.New()
		// 开启中间件
		e.Use(middleware.CORS())
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		// 加载路由
		routers.StepRouter(e)
		e.Logger.Fatal(e.Start(":46200"))
	}
}
