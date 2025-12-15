
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Ghaby-X/task_manager/Delivery/controllers"
	"github.com/Ghaby-X/task_manager/Delivery/routers"
	"github.com/Ghaby-X/task_manager/Infrastructure"
	"github.com/Ghaby-X/task_manager/Repositories"
	"github.com/Ghaby-X/task_manager/Usecases"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var Db *mongo.Database
var TasksCollection *mongo.Collection
var UsersCollection *mongo.Collection

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

	// Initialize services
	passwordService := Infrastructure.NewPasswordService()
	jwtService := Infrastructure.NewJWTService()

	// Initialize repositories
	taskRepo := Repositories.NewTaskRepository(TasksCollection)
	userRepo := Repositories.NewUserRepository(UsersCollection)

	// Initialize usecases
	taskUsecase := Usecases.NewTaskUsecase(taskRepo)
	userUsecase := Usecases.NewUserUsecase(userRepo, passwordService)

	// Initialize controller
	ctrl := controllers.NewController(taskUsecase, userUsecase, jwtService)

	// Initialize router
	r := routers.GetTaskRouter(ctrl, jwtService)
	err = r.Run()
	if err != nil {
		log.Fatal("failed to run router")
	}
}
