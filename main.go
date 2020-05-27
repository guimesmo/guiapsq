package main

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/handlers"
	"github.com/labstack/echo/v4"

	"github.com/guimesmo/guiapsq/repository/mongo"
)

var App *echo.Echo
var db *mongo.Connection

const (
	DBURL  = "mongodb://localhost:27017"
	DBName = "guiapsq"
)

func main() {
	// app setup
	App = echo.New()

	//database connection
	db = mongo.NewConnection(context.TODO(), DBURL, DBName)
	db.Connect()
	defer db.Disconnect()

	// first version routing
	h := &handlers.Handler{DB: db}
	App.GET("/api/psq/create", handlers.PsqCreate)

	// start app
	App.Start(":3232")

}
