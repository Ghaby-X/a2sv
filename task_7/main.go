package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Ghaby-X/task_manager/router"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var Db *mongo.Database
var TasksCollection *mongo.Collection
var UsersCollection *mongo.Collection

// dueDate := time.Now().Add(60 * time.Minute)
// seed = append(seed, models.Task{Title: "first task", Description: "sample task data", Status: "Pending", CreatedAt: time.Now(), DueDate: dueDate})
// seed = append(seed, models.Task{Title: "Second task", Description: "sample second task data", Status: "Approved", CreatedAt: time.Now(), DueDate: dueDate})
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not found in .env file")
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Successfully connected to MongoDB!")

	Db = client.Database(os.Getenv("DB_NAME"))
	TasksCollection = Db.Collection("tasks")
	UsersCollection = Db.Collection("users")

	r := router.GetTaskRouter(TasksCollection, UsersCollection) // Pass collections
	err = r.Run()
	if err != nil {
		log.Fatal("failed to run router")
	}
}