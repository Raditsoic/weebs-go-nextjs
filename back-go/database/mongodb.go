package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxPoolSize = 50
const timeoutSeconds = 30

var client *mongo.Client

func Connect() *mongo.Client {
	if client == nil {
		// Check if MONGO_URI is set, otherwise use the default connectionString
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			mongoURI = "mongodb://mongo:27017"
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		clientOptions = clientOptions.SetMaxPoolSize(maxPoolSize)

		ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
		defer cancel()

		var err error
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		// Ping the database to verify connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}

		log.Println("Connected to MongoDB at", mongoURI)
	}

	return client
}
