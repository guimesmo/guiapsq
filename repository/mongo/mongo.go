package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection implements the connection interface for MongoDB
type Connection struct {
	URL     string
	DBName  string
	context context.Context
	client  *mongo.Client
}

// NewConnection creates a new mongodb connection
func NewConnection(ctx context.Context, URL string, DBName string) Connection {
	if ctx == nil {
		ctx = context.TODO()
	}
	return Connection{URL: URL, DBName: DBName, context: ctx}

}

// Connect sets the db client connection
func (conn Connection) Connect() error {
	clientOptions := options.Client().ApplyURI(conn.URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	conn.client = client
	return err
}

// Disconnect disconnects the db instance
func (conn Connection) Disconnect() error {
	return conn.client.Disconnect(conn.context)
}

// GetCollection gets the db collection
func (conn Connection) GetCollection(name string) (*mongo.Collection, error) {
	collection := conn.client.Database(conn.DBName).Collection(name)
	return collection, nil
}
