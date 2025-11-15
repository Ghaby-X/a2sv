// Package models define the structure for data interface
package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Task defines the structure of the types
type Task struct {
	ID          int       `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DueDate     time.Time `json:"due_date,omitempty" bson:"due_date"`
}

var nextID = 3

func GetAllTasks() ([]Task, error) {
	var tasks []Task

	mongoCursor, err := MongoCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Printf("failed to extract all task")
		return nil, err
	}

	if err = mongoCursor.All(context.Background(), &tasks); err != nil {
		fmt.Println("failed to marshal tasks")
		return nil, err
	}

	return tasks, nil
}

// GetTaskByID : Return a task by id
func GetTaskByID(id int) (Task, error) {
	singleResult := MongoCollection.FindOne(context.TODO(), bson.M{"id": id})
	if singleResult.Err() != nil {
		fmt.Println("failed to extract single id")
		return Task{}, singleResult.Err()
	}

	var task Task
	err := singleResult.Decode(&task)
	if err != nil {
		fmt.Println("failed to marshal task")
		fmt.Println(err)
		return Task{}, err
	}

	return task, nil
}

// UpdateTask : updata tasks in database
func UpdateTask(id int, newTask Task) error {
	var task Task
	err := MongoCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		fmt.Println(err)
		return err
	}
	newTask.ID = id

	// check for zeroed out input
	if newTask.Title == "" {
		newTask.Title = task.Title
	}

	if newTask.Description == "" {
		newTask.Description = task.Description
	}

	if newTask.Status == "" {
		newTask.Status = task.Status
	}

	if newTask.DueDate.IsZero() {
		newTask.DueDate = task.DueDate
	}

	if newTask.CreatedAt.IsZero() {
		newTask.CreatedAt = task.CreatedAt // preserve CreatedAt
	}

	_, err = MongoCollection.ReplaceOne(context.TODO(), bson.M{"id": id}, newTask)
	if err != nil {
		fmt.Printf("failed to update task: ")
		fmt.Println(err)
		return err
	}

	return nil
}

// DeleteTasks : remove tasks from database
func DeleteTasks(id int) error {
	_, err = MongoCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		fmt.Println("failed to delete task")
		return err
	}

	return nil
}

// CreateTask : creates a new task
func CreateTask(task Task) error {
	task.ID = nextID
	task.CreatedAt = time.Now()
	nextID += 1

	bsonData, err := bson.Marshal(task)
	if err != nil {
		fmt.Println("failed to marshal bson data")
		return err
	}

	insertOneResult, err := MongoCollection.InsertOne(context.TODO(), bsonData)
	if err != nil {
		fmt.Println("failed to insert task into MongoClient")
		return err
	}
	fmt.Println(insertOneResult)
	return nil
}
