// Package models define the structure for data interface
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task defines the structure of the types
type Task struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	DueDate     time.Time          `json:"due_date,omitempty" bson:"due_date,omitempty"`
}

type CreateTaskDTO struct {
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	DueDate     time.Time `json:"due_date,omitempty" bson:"due_date,omitempty"`
}
