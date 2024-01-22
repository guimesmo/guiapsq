package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/guimesmo/guiapsq/repository"
)

type Client struct {
	mongo
}

const ClientCollectionName = "client"

func NewClientRepository(s *mgo.Session, db string) repository.User {
	return &Client{
		mongo{
			session:        s,
			db:             db,
			collectionName: ClientCollectionName,
		},
	}
}

func (c Client) Find() {

}
func (c Client) Get() {

}
func (c Client) Delete() error {
	return nil
}
