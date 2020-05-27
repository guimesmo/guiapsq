package services

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/models"
)

// PsqCreate is the service to create a Psq instance
func PsqCreate(c context.Context) (models.PsqImp, error) {
	collection, _ := db.GetCollection(c, "psq")
	psq := new(models.PsqImp)

	if err := c.Bind(psq); err != nil {
		return nil, err
	}
	item, err := collection.InsertOne(context.TODO(), author)

	return item, err
}
