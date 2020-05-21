package mongo

import (
	"context"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection implements the connection interface for MongoDB
type Connection struct {
	URL        string
	DBName     string
	context context
	client *mongo.Client
}

func NewConnectionn(URL string, DBName string, ctx context) Connection {
	if ctx == nil {
		ctx = context.TODO()
	}
	return new(Connection{URL: URL, DBName: DBName})

} 

// Connect sets the db client connection
func (conn Connection) Connect() error {
	clientOptions := options.Client().ApplyURI(conn.URL)
	conn.client, err := mongo.Connect(context.TODO(), clientOptions)
	return err
}

// GetCollection gets the db collection
func (conn Connection) GetCollection(name string) (*mongo.Collection, error) {
	collection := conn.client.Database(conn.client.DBName).Collection(name)
	return collection, nil
}

func (conn Connection) Disconect() error {
	conn.client.Disconnect(context.TODO())
}