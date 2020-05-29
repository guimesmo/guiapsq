package services

import (
	"github.com/guimesmo/guiapsq/internal/psq/models"
	"github.com/guimesmo/guiapsq/repository/mongo"
	"github.com/labstack/echo/v4"
)

// PsqCreate is the service to create a Psq instance
func PsqCreate(c echo.Context, conn *mongo.Connection) (models.Psq, error) {
	collection, _ := conn.GetCollection("psq")
	psq := new(models.PsqImp)

	if err := c.Bind(psq); err != nil {
		return nil, err
	}
	return collection.InsertOne(c.Request().Context(), psq)
}
