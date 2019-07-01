package test

import (
	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	return c.String(200, "test")
}