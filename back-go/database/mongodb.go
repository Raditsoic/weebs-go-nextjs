package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb://localhost:27017/"

const databaseName string = "weebs"
const maxPoolSize = 50
const timeoutSeconds = 30

// Connect establishes a connection to the MongoRepo database.
func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions = clientOptions.SetMaxPoolSize(maxPoolSize)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
