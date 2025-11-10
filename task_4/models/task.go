// Package models define the structure for data interface
package models

import (
	"time"
)

// Tasks defines the structure of the types
type Tasks struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `created_at:"time"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `due_date:"string"`
}
