package main

import (
	handlers "github.com/guimesmo/guiapsq/internal/app/psq/handler"
	"github.com/labstack/echo/v4"
)

const (
	DBURL  = "mongodb://localhost:27017"
	DBName = "guiapsq"
)

func main() {
	// app setup
	app := echo.New()
	h := handlers.NewClinetHandler()
	handlers.Route(app, h)

}
