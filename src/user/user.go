package user

import (
	"time"
	util "web_model/src/util"
	models "web_model/src/model"
	db "web_model/src/db"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// 用户注册
func Register(c echo.Context) error {
	// 获取注册参数
	name := c.FormValue("name")
	password := c.FormValue("password")
	// 判断用户名是否重复
	user := GetUserByName(name)
	if user.Name != "" {
		return c.JSON(200, util.ReturnBody(1, "", "用户名已存在，请换个名称重试！"))
	}
	// 密码加密
	salt := time.Now().String()
	encryptPwd := util.EncryptStr(salt, password)
	// 新增用户
	err := db.Create("user", models.User{"", name, encryptPwd, salt})
	if err != nil {
		panic(err)
		return c.JSON(200, util.ReturnBody(1, "", "注册失败，发生异常！"))
	}
	return c.JSON(200, util.ReturnBody(0, "", "注册成功"))
}

// 用户登录
func Login(c echo.Context) error {
	// 获取参数
	name := c.FormValue("name")
	password := c.FormValue("password")
	// 判断用户名是否存在
	user := GetUserByName(name)
	if user.Name == "" {
		return c.JSON(200, util.ReturnBody(1, "", "用户不存在！"))
	}
	// 判断密码是否一致
	encryptPwd := util.EncryptStr(user.Salt, password)
	if encryptPwd == user.Password {
		token := MakeToken(name, user.Id)
		return c.JSON(200, util.ReturnBody(0, token, "登录成功"))
	} else {
		return c.JSON(200, util.ReturnBody(1, "", "登录失败,密码错误！"))
	}
}

// 根据用户名获取用户信息
func GetUserByName(name string) models.User {
	user := models.User{}
	db.FindOne("user", &user, &bson.M{"name": name}, nil)
	return user
}