
// Package controllers interacts with request and handle http related processing
package controllers

import (
	"net/http"

	"github.com/Ghaby-X/task_manager/Domain"
	"github.com/Ghaby-X/task_manager/Infrastructure"
	"github.com/Ghaby-X/task_manager/Usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller struct to hold usecases
type Controller struct {
	taskUsecase Usecases.TaskUsecase
	userUsecase Usecases.UserUsecase
	jwtService  Infrastructure.JWTService
}

// NewController creates a new Controller instance
func NewController(taskUsecase Usecases.TaskUsecase, userUsecase Usecases.UserUsecase, jwtService Infrastructure.JWTService) *Controller {
	return &Controller{
		taskUsecase: taskUsecase,
		userUsecase: userUsecase,
		jwtService:  jwtService,
	}
}

// HandleGetTask to handle get tasks
func (ctrl *Controller) HandleGetTask(c *gin.Context) {
	tasks, err := ctrl.taskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to get task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// HandleCreateTask create tasks
func (ctrl *Controller) HandleCreateTask(c *gin.Context) {
	var newTask Domain.Task

	err := c.BindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad request"})
		return
	}

	err = ctrl.taskUsecase.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "task created successfully"})
}

// HandleGetTaskByID gets a task by id
func (ctrl *Controller) HandleGetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid task ID format"})
		return
	}

	task, err := ctrl.taskUsecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (ctrl *Controller) HandleDeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid task ID format"})
		return
	}

	err = ctrl.taskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "task deleted successfully"})
}

func (ctrl *Controller) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid task ID format"})
		return
	}

	var newTask Domain.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to parse task json"})
		return
	}

	err = ctrl.taskUsecase.UpdateTask(id, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "task updated successfully"})
}

// HandleRegister registers a new user
func (ctrl *Controller) HandleRegister(c *gin.Context) {
	var user Domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.userUsecase.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// HandleLogin authenticates a user and returns a JWT
func (ctrl *Controller) HandleLogin(c *gin.Context) {
	var userInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userUsecase.Login(userInput.Username, userInput.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := ctrl.jwtService.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// HandlePromoteUser promotes a user to admin
func (ctrl *Controller) HandlePromoteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid user ID format"})
		return
	}

	if err := ctrl.userUsecase.PromoteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}
