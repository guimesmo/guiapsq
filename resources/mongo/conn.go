package conn


import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetClient(c echo.Context) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func GetCollection(c echo.Context, name string) (*mongo.Collection, error) {
	client, _ := GetClient(c)
	collection := client.Database("guiapsq").Collection(name)
	return collection, nil
}
