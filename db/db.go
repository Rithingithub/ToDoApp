package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // always remember to cancel the context to avoid leaking resources

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB server: %w", err)
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

// GetDB returns a reference to the 'test' database
func GetDB() *mongo.Database {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("test")
}
