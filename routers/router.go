package routers

import (
	"web_model/middleware/jwt"
	v1 "web_model/routers/api/v1"

	"github.com/labstack/echo/v4"
)

func StepRouter(e *echo.Echo) {

	// 用户登录
	e.POST("/user/login", v1.Login)

	// 用户注册
	e.POST("/user/register", v1.Register)

	// 路由分组
	apiV1 := e.Group("/api/v1")
	apiV1.Use(jwt.ValidateToken)
	{
		apiV1.GET("/test", v1.TestAuth)
	}

}
