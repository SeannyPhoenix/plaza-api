package db

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient     *mongo.Client
	database        *mongo.Database
	collections     map[string]*mongo.Collection = make(map[string]*mongo.Collection)
	registryBuilder *bsoncodec.RegistryBuilder   = bson.NewRegistryBuilder()
)

// Connect to the db
func Connect(
	ctx context.Context,
	mongoURL string,
	databaseName string,
) {
	var err error
	if mongoClient, err = mongo.NewClient(
		options.Client().ApplyURI(mongoURL),
		options.Client().SetRegistry(registryBuilder.Build()),
	); err != nil {
		log.Fatal(err)
	}

	if err = mongoClient.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	database = mongoClient.Database(databaseName)
}

// Disconnect from the db
func Disconnect(ctx context.Context) {
	mongoClient.Disconnect(ctx)
	database = nil
	collections = make(map[string]*mongo.Collection)
}

// IsConnected returns the connection status
func IsConnected() bool {
	return database != nil
}

// collection retrieves an instance of a collection
func collection(collectionName string) *mongo.Collection {
	if coll, ok := collections[collectionName]; ok {
		return coll
	}
	collections[collectionName] = database.Collection(collectionName)
	return collections[collectionName]
}

func checkDup(err error) (bool, string) {
	writeEx, ok := err.(mongo.WriteException)
	if !ok {
		return false, ""
	}
	for _, writeErr := range writeEx.WriteErrors {
		if writeErr.Code == 11000 {
			msg := writeErr.Message
			if idx := strings.Index(msg, " index: "); idx >= 0 {
				msg = msg[idx+8:]
				if idx = strings.IndexRune(msg, ' '); idx >= 0 {
					return true, msg[:idx]
				}
			}
		}
	}
	return false, ""
}
