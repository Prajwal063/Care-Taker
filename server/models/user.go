package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents the schema for the User collection in MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" validate:"required,min=3,max=32"`
	Email     string             `bson:"email" validate:"required,email"`
	GoogleID  string             `bson:"googleId" validate:"required"`
	Picture   string             `bson:"picture" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}

// CreateIndexes creates the necessary indexes for the User collection
func CreateIndexes(collection *mongo.Collection) error {
	// Define the compound index with multiple keys
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},    // 1 for ascending order
			{Key: "googleId", Value: 1}, // 1 for ascending order
		},
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}

// ToJSON removes MongoDB specific fields and returns a clean JSON representation of the User
func (u *User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID.Hex(),
		"name":     u.Name,
		"email":    u.Email,
		"googleId": u.GoogleID,
		"picture":  u.Picture,
	}
}
