package user

import (
	"time"
	"fmt"
	util "web_model/src/util"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// 验证token
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("token")
		if tokenStr == "" {
			return c.JSON(200, util.ReturnBody(1, "", "token不能为空！"))
		}
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return c.JSON(200, util.ReturnBody(1, "", "token解析失败"))
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["id"]
			expDate := claims["expDate"]
			timeNow := time.Now().Format("2006-01-02 15:04:05")
			c.Request().Header.Set("userId", userId.(string))
			if (timeNow > expDate.(string)) {
				return c.JSON(200, util.ReturnBody(1, "", "token已失效，请重新获取"))
			}
			if userId == "" {
				return c.JSON(200, util.ReturnBody(1, "", "token不合法"))
			}
		} else {
			return c.JSON(200, util.ReturnBody(1, "", "token解析失败"))
		}
		return next(c)
	}
}

// 生成Token
func MakeToken(name string, id bson.ObjectId) string {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["id"] = id.Hex()
	claims["expDate"] = time.Now().Add(time.Hour * 1).Format("2006-01-02 15:04:05")
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	return t
}
