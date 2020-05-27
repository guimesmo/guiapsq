package routes

import (
	"github.com/guimesmo/guiapsq/internal/psq/handlers"
	"github.com/labstack/echo/v4"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.POST("/api/psq/create", handlers.PsqCreate)
}

// criar um serviço e injetar nele os handler
