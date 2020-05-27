package services

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/models"
	"github.com/guimesmo/guiapsq/repository"
	"github.com/labstack/echo/v4"
)

// PsqCreate is the service to create a Psq instance
func PsqCreate(c echo.Context, conn repository.Connection) (models.PsqImp, error) {
	collection := conn.GetCollection("psq")
	psq := new(models.PsqImp)

	if err := c.Bind(psq); err != nil {
		return nil, err
	}
	item, err := collection.InsertOne(context.TODO(), psq)

	return item, err
}
