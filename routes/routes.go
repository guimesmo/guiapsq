package routes

import (
	"github.com/labstack/echo/v4"
)

var App *echo.Echo

func init() {
	App = echo.New()
}
