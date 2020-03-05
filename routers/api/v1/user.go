package v1

import (
	"web_model/middleware/jwt"
	models "web_model/models"
	"web_model/pkg/e"
	"web_model/pkg/encrypt"
	"web_model/service/user_service"

	"github.com/labstack/echo/v4"
)

// 用户注册
func Register(c echo.Context) error {

	// 参数绑定
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(200, e.ReturnBody(1, "", "参数绑定失败"))
	}

	// 判断用户名是否重复
	userSvc := &user_service.UserSvc{}
	dbUser, err := userSvc.GetByName(user.Name)
	if err != nil {
		return c.JSON(200, e.ReturnBody(1, "", "注册失败"+err.Error()))
	}
	if dbUser.Name != "" {
		return c.JSON(200, e.ReturnBody(1, "", "用户名已存在，请换个名称重试！"))
	}

	// 新增用户
	if err = userSvc.Add(user); err != nil {
		return c.JSON(200, e.ReturnBody(1, "", "注册失败"+err.Error()))
	}

	return c.JSON(200, e.ReturnBody(0, "", "注册成功"))

}

// 用户登录
func Login(c echo.Context) error {

	// 绑定数据
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(200, e.ReturnBody(1, "", "参数绑定失败"))
	}

	// 查询用户
	userSvc := &user_service.UserSvc{}
	dbUser, err := userSvc.GetByName(user.Name)
	if err != nil {
		return c.JSON(200, e.ReturnBody(1, "", "登录失败"))
	}
	if dbUser.Name == "" {
		return c.JSON(200, e.ReturnBody(1, "", "用户不存在！"))
	}

	// 判断密码是否一致
	ePassword := encrypt.Md5WithSalt(user.Password, dbUser.Salt)
	if ePassword == dbUser.Password {
		token := jwt.MakeToken(user.Name, dbUser.Id)
		return c.JSON(200, e.ReturnBody(0, token, "登录成功"))
	} else {
		return c.JSON(200, e.ReturnBody(1, "", "登录失败,密码错误！"))
	}

}
