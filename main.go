package main


import (
	"github.com/labstack/echo/v4"
	"github.com/guimesmo/guiapsq/controllers"
)

var App *echo.Echo

func main() {
	App = echo.New()

	App.POST("/api/psq/create", controllers.PsqController)
	routes.App.Start(":3232")
}
