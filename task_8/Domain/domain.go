
// Package Domain defines the structure for data interface
package Domain

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

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-" bson:"password"`
	Role     string             `json:"role" bson:"role"` // e.g., "admin", "user"
}
