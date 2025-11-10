// Package models for holding models
package models

import "sync"

type Book struct {
	ID         int
	Title      string
	Author     string
	Status     string
	ReservedBy int
	Mutex      sync.Mutex
}
