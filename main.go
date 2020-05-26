package main

import (
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
	db = new(mongo.NewConnection{URL: DBURL, DBName: DBName, context: App.GetContext()})
	db.Connect()
	App.GetContext().Set("db", db)
	defer db.Disconnect()

	// first version routing
	App.GET("/api/psq/create", handlers.PsqCreate)

	// start app
	App.Start(":3232")

}
