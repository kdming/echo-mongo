package test

import (
	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	userId := c.Request().Header.Get("userId")
	return c.String(200, userId)
}
