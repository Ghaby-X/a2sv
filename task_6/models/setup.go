package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient     *mongo.Client
	MongoCollection *mongo.Collection
	err             error
)

const (
	dbName         = "a2sv"
	collectionName = "task"
)

// SetupMongoDBClient starsts mongo db
func SetupMongoDBClient() error {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("failed to load mongo db uri from env")
	}

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("failed to create mongo client")
	}

	// verify connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("mongo db not live - ping not successful")
		return err
	}

	// setting up connection
	MongoCollection = mongoClient.Database(dbName).Collection(collectionName)
	return nil
}
