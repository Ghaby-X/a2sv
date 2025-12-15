package data

import (
	"context"
	"fmt"
	"time"

	"github.com/Ghaby-X/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" // Import mongo package
)

// GetAllTasks gets all task in the database
func GetAllTasks(tasksCollection *mongo.Collection) ([]models.Task, error) {
	var tasks []models.Task

	mongoCursor, err := tasksCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Printf("failed to extract all task")
		return nil, err
	}

	// marshal bson data into tasks array
	if err = mongoCursor.All(context.Background(), &tasks); err != nil {
		fmt.Println("failed to marshal tasks")
		return nil, err
	}

	return tasks, nil
}

// GetTaskByID : Return a task by id
func GetTaskByID(id primitive.ObjectID, tasksCollection *mongo.Collection) (models.Task, error) {
	singleResult := tasksCollection.FindOne(context.TODO(), bson.M{"_id": id})
	if singleResult.Err() != nil {
		fmt.Println("failed to extract single id")
		return models.Task{}, singleResult.Err()
	}

	var task models.Task
	err := singleResult.Decode(&task)
	if err != nil {
		fmt.Println("failed to marshal task")
		fmt.Println(err)
		return models.Task{}, err
	}

	return task, nil
}

// UpdateTask : updata tasks in database
func UpdateTask(id primitive.ObjectID, newTask models.Task, tasksCollection *mongo.Collection) error {
	var task models.Task
	err := tasksCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		fmt.Println(err)
		return err
	}
	newTask.ID = task.ID

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

	_, err = tasksCollection.ReplaceOne(context.TODO(), bson.M{"_id": id}, newTask)
	if err != nil {
		fmt.Printf("failed to update task: ")
		fmt.Println(err)
		return err
	}

	return nil
}

// DeleteTasks : remove tasks from database
func DeleteTasks(id primitive.ObjectID, tasksCollection *mongo.Collection) error {
	_, err := tasksCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println("failed to delete task")
		return err
	}

	return nil
}

// CreateTask : creates a new task
func CreateTask(task models.Task, tasksCollection *mongo.Collection) error {
	task.CreatedAt = time.Now()
	task.ID = primitive.NewObjectID() // Assign a new ObjectID here

	_, err := tasksCollection.InsertOne(context.TODO(), task)
	if err != nil {
		fmt.Println("failed to insert task into MongoClient")
		return err
	}
	// fmt.Println(insertOneResult) // This was uncommented, but insertOneResult is not used.
	return nil
}