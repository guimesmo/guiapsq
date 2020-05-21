package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/guimesmo/guiapsq/internal/app/psq/handlers"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.POST("/api/psq/create", handlers.PsqCreate)
}


//#mudar controller pra handler
// no db criar um struct com a conexão
// criar um serviço e injetar nele os handler