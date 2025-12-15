
package Usecases

import (
	"time"

	"github.com/Ghaby-X/task_manager/Domain"
	"github.com/Ghaby-X/task_manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase interface {
	GetAllTasks() ([]Domain.Task, error)
	GetTaskByID(id primitive.ObjectID) (Domain.Task, error)
	UpdateTask(id primitive.ObjectID, newTask Domain.Task) error
	DeleteTask(id primitive.ObjectID) error
	CreateTask(task Domain.Task) error
}

type taskUsecase struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUsecase(taskRepo Repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) GetAllTasks() ([]Domain.Task, error) {
	return u.taskRepo.GetAllTasks()
}

func (u *taskUsecase) GetTaskByID(id primitive.ObjectID) (Domain.Task, error) {
	return u.taskRepo.GetTaskByID(id)
}

func (u *taskUsecase) UpdateTask(id primitive.ObjectID, newTask Domain.Task) error {
	// check for zeroed out input
	existingTask, err := u.taskRepo.GetTaskByID(id)
	if err != nil {
		return err
	}

	if newTask.Title == "" {
		newTask.Title = existingTask.Title
	}

	if newTask.Description == "" {
		newTask.Description = existingTask.Description
	}

	if newTask.Status == "" {
		newTask.Status = existingTask.Status
	}

	if newTask.DueDate.IsZero() {
		newTask.DueDate = existingTask.DueDate
	}

	newTask.CreatedAt = existingTask.CreatedAt 

	return u.taskRepo.UpdateTask(id, newTask)
}

func (u *taskUsecase) DeleteTask(id primitive.ObjectID) error {
	return u.taskRepo.DeleteTasks(id)
}

func (u *taskUsecase) CreateTask(task Domain.Task) error {
	task.CreatedAt = time.Now()
	return u.taskRepo.CreateTask(task)
}
