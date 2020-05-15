package routes

import (
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()
}