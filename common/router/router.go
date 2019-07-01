package router

import (
	user "web_model/src/user"
	test "web_model/src/test"
	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {

	// 用户登录
	e.POST("/user/login", user.Login)
	// 用户注册
	e.POST("/user/register", user.Register)
	
	// 路由分组
	api := e.Group("/api")
	// 开启token验证
	api.Use(user.ValidateToken);
	// 测试方法
	api.GET("/test", test.Test)

}