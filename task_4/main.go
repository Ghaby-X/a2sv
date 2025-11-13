package main

import (
	"fmt"
	"time"

	"github.com/Ghaby-X/task_manager/models"
	"github.com/Ghaby-X/task_manager/router"
)

func SeedTasks() {
	seed := []models.Task{}
	dueDate := time.Now().Add(60 * time.Minute)
	seed = append(seed, models.Task{ID: 1, Title: "first task", Description: "sample task data", Status: "Pending", CreatedAt: time.Now(), DueDate: dueDate})
	seed = append(seed, models.Task{ID: 2, Title: "Second task", Description: "sample second task data", Status: "Approved", CreatedAt: time.Now(), DueDate: dueDate})

	models.Tasks = seed
}

func main() {
	SeedTasks()

	r := router.GetTaskRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("Failed to run error")
	}
}
