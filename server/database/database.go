package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client
var ServiceCollection *mongo.Collection
var EventCollection *mongo.Collection

func OpenCollection(collectionName string) *mongo.Collection {
	return Client.Database("care-taker").Collection(collectionName)
}
