package controllers

import (
	"context",
	"net/http"
	"github.com/labstack/echo"
	db "github.com/guimesmo/guiapsq/resources/mongo"
	"github.com/guimesmo/guiapsq/internal/app/psq/models"
)

func PsqCreate(c echo.Context) error {
	collection = db.GetCollection(c, "psq")
	psq = new(models.psq)
	if err := c.Bind(psq); err != nil {
		return err
	}
	item, err := collection.InsertOne(context.TODO(), psq)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)

}