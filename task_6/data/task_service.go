// Package services contains the essential services
package services

import (
	"github.com/Ghaby-X/task_manager/models"
)

func GetTasks() ([]models.Task, error) {
	return models.GetAllTasks()
}

func GetTaskByID(id int) (models.Task, error) {
	return models.GetTaskByID(id)
}

func CreateTask(task models.Task) error {
	return models.CreateTask(task)
}

func UpdateTask(id int, task models.Task) error {
	return models.UpdateTask(id, task)
}

func DeleteTask(taskID int) error {
	return models.DeleteTasks(taskID)
}
