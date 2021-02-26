package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var database *mongo.Database
var accountCollection *mongo.Collection

// Connect to the db
func Connect(mongoURL string, databaseName string) *mongo.Client {
	var err error
	if mongoClient, err = mongo.NewClient(
		options.Client().ApplyURI(mongoURL),
	); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database = mongoClient.Database(databaseName)
	accountCollection = database.Collection("accounts")

	log.Println("database connected")
	return mongoClient
}
