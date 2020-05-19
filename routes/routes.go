package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/guimesmo/guiapsq/internal/app/psq/controllers"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.POST("/api/psq/create", controllers.PsqController)
}
