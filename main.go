package main


import (
	"github.com/labstack/echo/v4"
	"github.com/guimesmo/guiapsq/internal/app/psq/handlers"

	"github.com/guimesmo/guiapsq/repository/mongo"
)

var App *echo.Echo
var db *mongo.Connection

const (
	DBURL := "mongodb://localhost:27017"
	DBName := "guiapsq"
)


func main() {
	// app setup
	App = echo.New()
	ctx := App.AcquireContext()
	defer ctx.ReleaseContext()
	
	//database connection
	db = new(mongo.Connection{URL: DBURL, DBName: DBName, context: ctx})
	db.Connect()
	defer db.Disconnect(ctx)

	// first version routing
	App.GET("/api/psq/create", handlers.PsqCreate)

	// start app
	App.Start(":3232")


}
