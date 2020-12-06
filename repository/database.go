package repository

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/models"
)


type PsqRepository interface {
	Insert(context.Context, models.Psq)
	Find(context.Context)
	Delete()
	Update()
}
