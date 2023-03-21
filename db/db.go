package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()               // always remember to cancel the context to avoid leaking resources
	defer client.Disconnect(ctx) // always remember to disconnect the client to avoid leaking resources

	col := client.Database("test").Collection("users")
	_ = col // to avoid a "declared but not used" error
}
