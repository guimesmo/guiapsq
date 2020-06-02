package main

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/handlers"
	"github.com/guimesmo/guiapsq/internal/psq/services"
	"github.com/labstack/echo/v4"

	"github.com/guimesmo/guiapsq/repository/mongo"
)

// App the main running app
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
	db = mongo.NewConnection(context.Background(), DBURL, DBName)
	db.Connect()
	defer db.Disconnect(context.TODO())

	psqRepository := mongo.NewPsq(db)
	psqService := services.NewPsq(psqRepository)

	// first version routing
	h := handlers.NewHandler(psqService)
	App.GET("/api/psq/create", h.PsqCreate)

	// start app
	App.Start(":3232")

}
