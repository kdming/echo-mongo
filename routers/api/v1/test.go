package v1

import (
	"web_model/pkg/e"

	"github.com/labstack/echo/v4"
)

func TestAuth(c echo.Context) error {
	return c.JSON(200, e.ReturnBody(0, "ok", "验证成功"))
}
