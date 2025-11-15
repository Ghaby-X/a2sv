package main

import (
	"log"
	"time"

	"github.com/Ghaby-X/task_manager/models"
	"github.com/Ghaby-X/task_manager/router"
	"github.com/joho/godotenv"
)

func SeedTasks() {
	seed := []models.Task{}
	dueDate := time.Now().Add(60 * time.Minute)
	seed = append(seed, models.Task{ID: 1, Title: "first task", Description: "sample task data", Status: "Pending", CreatedAt: time.Now(), DueDate: dueDate})
	seed = append(seed, models.Task{ID: 2, Title: "Second task", Description: "sample second task data", Status: "Approved", CreatedAt: time.Now(), DueDate: dueDate})

	for _, task := range seed {
		models.CreateTask(task)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup mongo db
	err = models.SetupMongoDBClient()
	if err != nil {
		log.Fatal("failed to setup mongodb")
	}

	SeedTasks()

	r := router.GetTaskRouter()
	err = r.Run()
	if err != nil {
		log.Fatal("failed to run router")
	}
}
