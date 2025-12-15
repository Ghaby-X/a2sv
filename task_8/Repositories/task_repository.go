
package Repositories

import (
	"context"
	"fmt"

	"github.com/Ghaby-X/task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	GetAllTasks() ([]Domain.Task, error)
	GetTaskByID(id primitive.ObjectID) (Domain.Task, error)
	UpdateTask(id primitive.ObjectID, newTask Domain.Task) error
	DeleteTasks(id primitive.ObjectID) error
	CreateTask(task Domain.Task) error
}

type taskRepository struct {
	tasksCollection *mongo.Collection
}

func NewTaskRepository(tasksCollection *mongo.Collection) TaskRepository {
	return &taskRepository{tasksCollection: tasksCollection}
}

func (r *taskRepository) GetAllTasks() ([]Domain.Task, error) {
	var tasks []Domain.Task

	mongoCursor, err := r.tasksCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("failed to find tasks: %w", err)
	}

	if err = mongoCursor.All(context.Background(), &tasks); err != nil {
		return nil, fmt.Errorf("failed to decode tasks: %w", err)
	}

	return tasks, nil
}

func (r *taskRepository) GetTaskByID(id primitive.ObjectID) (Domain.Task, error) {
	var task Domain.Task
	filter := bson.M{"_id": id}
	err := r.tasksCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return Domain.Task{}, fmt.Errorf("task not found: %w", err)
	}
	return task, nil
}

func (r *taskRepository) UpdateTask(id primitive.ObjectID, newTask Domain.Task) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": newTask}
	_, err := r.tasksCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (r *taskRepository) DeleteTasks(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.tasksCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (r *taskRepository) CreateTask(task Domain.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := r.tasksCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}
