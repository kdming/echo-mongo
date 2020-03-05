package jwt

import (
	"fmt"
	"time"
	"web_model/pkg/e"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

// 验证token
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获取token字符串
		tokenStr := c.Request().Header.Get("token")
		if tokenStr == "" {
			return c.JSON(200, e.ReturnBody(1, "", "token不能为空！"))
		}
		// 解密token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return c.JSON(200, e.ReturnBody(1, "", "token解析失败"))
		}
		// 获取token内容
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["userId"]
			expDate := claims["expDate"]
			timeNow := time.Now().Format("2006-01-02 15:04:05")
			if timeNow > expDate.(string) {
				return c.JSON(200, e.ReturnBody(1, "", "token已失效，请重新获取"))
			}
			if userId == "" {
				return c.JSON(200, e.ReturnBody(1, "", "token不合法"))
			}
			// 将用户id放入请求头
			c.Request().Header.Set("userId", userId.(string))
		} else {
			return c.JSON(200, e.ReturnBody(1, "", "token解析失败"))
		}

		return next(c)
	}
}

// 生成Token
func MakeToken(name string, id bson.ObjectId) string {
	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)
	// 绑定数据
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["userId"] = id.Hex()
	claims["expDate"] = time.Now().Add(time.Hour * 1).Format("2006-01-02 15:04:05")
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	return t
}
