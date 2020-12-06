package mongo

import (
	"context"

	"github.com/guimesmo/guiapsq/internal/psq/models"
	"github.com/guimesmo/guiapsq/repository"
)

type Psq struct {
	conn *Connection
}

func NewPsq(conn *Connection) repository.PsqRepository {
	return &Psq{conn}
}

func (p Psq) Insert(c context.Context, psq models.Psq) {
	collection, _ := p.conn.GetCollection("psq")
	collection.InsertOne(c, psq)
}

func (p Psq) Find(context.Context) {}
func (p Psq) Delete()              {}
func (p Psq) Update()              {}
