package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"care-taker/database"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Database() *mongo.Client {
	MONGO_URL := os.Getenv("MONGO_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database!")

	return client
}

func Collections() {
	database.ServiceCollection = database.OpenCollection("services")
	database.EventCollection = database.OpenCollection("events")
	database.UserCollection = database.OpenCollection("users")
}
