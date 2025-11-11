// Package models define the structure for data interface
package models

import (
	"errors"
	"time"
)

// Task defines the structure of the types
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date,omitempty"`
}

var (
	Tasks  []Task
	nextID = 1
)

func GetAllTasks() []Task {
	return Tasks
}

// GetTaskByID : Return a task by id
func GetTaskByID(id int) (Task, error) {
	for _, task := range Tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return Task{}, errors.New("Task Not found")
}

// UpdateTask : updata tasks in database
func UpdateTask(id int, newTask Task) error {
	for i, task := range Tasks {
		if task.ID == id {
			newTask.ID = id
			if newTask.CreatedAt.IsZero() {
				newTask.CreatedAt = task.CreatedAt // preserve CreatedAt
			}
			Tasks[i] = newTask
			return nil
		}
	}

	return errors.New("task not found")
}

// DeleteTasks : remove tasks from database
func DeleteTasks(id int) error {
	newTasks := []Task{}

	for _, task := range Tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	Tasks = newTasks
	return nil
}

// CreateTask : creates a new task
func CreateTask(task Task) error {
	task.ID = nextID
	task.CreatedAt = time.Now()
	Tasks = append(Tasks, task)
	nextID += 1
	return nil
}
