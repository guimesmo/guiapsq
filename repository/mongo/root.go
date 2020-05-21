package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Creatable interface {
	SetID(ID string)
	SetCreated(t time.Time)
}

type mongo struct {
	session        *mgo.Session
	db             string
	collectionName string
}

func (m *mongo) getCollection() (*mgo.Collection, func()) {
	session := m.session.Copy()
	return session.DB(m.db).C(m.collectionName), session.Close
}

func (m *mongo) findByFilter(filter interface{}, result interface{}) error {
	collection, close := m.getCollection()
	defer close()

	return collection.Find(filter).One(result)
}

func (m *mongo) findAllByFilter(filter interface{}, result []interface{}) error {
	collection, close := m.getCollection()
	defer close()

	return collection.Find(filter).All(result)

}

func (m *mongo) insert(data interface{}) error {
	collection, close := m.getCollection()
	defer close()

	if d, ok := data.(Creatable); ok {
		d.SetID(bson.NewObjectId().Hex())
		d.SetCreated(time.Now())
	}

	return collection.Insert(data)
}
