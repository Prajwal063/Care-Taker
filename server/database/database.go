package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go-jwt").Collection(collectionName)
}
