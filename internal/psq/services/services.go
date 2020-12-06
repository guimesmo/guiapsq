package services

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/models"
	"github.com/guimesmo/guiapsq/repository"
)

type PsqService interface {
	Insert(context.Context, models.Psq)
	Delete()
	Update()
}

type Psq struct {
	db repository.PsqRepository
}

func NewPsq(db repository.PsqRepository) PsqService {
	return Psq{db}
}

// Inser is the service to create a Psq instance
func (p Psq) Insert(c context.Context, instance models.Psq) {
	p.db.Insert(c, instance)
}

func (p Psq) Delete() {}

func (p Psq) Update() {}
