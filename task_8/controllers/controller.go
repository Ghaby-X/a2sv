// Package controllers interacts with request and handle http related processing
package controllers

import (
	"fmt"
	"net/http"

	"github.com/Ghaby-X/task_manager/models"
	"github.com/Ghaby-X/task_manager/data"
	"github.com/Ghaby-X/task_manager/middleware" // Import the middleware package
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" // New import
)

// Controller struct to hold database collections
type Controller struct {
	tasksCollection *mongo.Collection
	usersCollection *mongo.Collection
}

// NewController creates a new Controller instance
func NewController(tasksCol, usersCol *mongo.Collection) *Controller {
	return &Controller{
		tasksCollection: tasksCol,
		usersCollection: usersCol,
	}
}

// HandleGetTask to handle get tasks
func (ctrl *Controller) HandleGetTask(c *gin.Context) {
	tasks, err := data.GetAllTasks(ctrl.tasksCollection)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to get task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// HandleCreateTask create tasks
func (ctrl *Controller) HandleCreateTask(c *gin.Context) {
	var newTask models.Task

	err := c.BindJSON(&newTask)
	if err != nil {
		fmt.Println("Failed to bind json")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Bad request"})
		return
	}

	// print out new task
	fmt.Println("task retrieved successfully")
	err = data.CreateTask(newTask, ctrl.tasksCollection)
	if err != nil {
		println("failed to create task")
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

	task, err := data.GetTaskByID(id, ctrl.tasksCollection)
	if err != nil {
		fmt.Println("task not found, id: ", idParam)
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

	err = data.DeleteTasks(id, ctrl.tasksCollection)
	if err != nil {
		fmt.Println("failed to delete task")
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

	// update task
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		fmt.Println("failed to parse task")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to parse task json"})
		return
	}

	err = data.UpdateTask(id, newTask, ctrl.tasksCollection)
	if err != nil {
		fmt.Println("failed to update task")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "task updated successfully"})
}

// HandleRegister registers a new user
func (ctrl *Controller) HandleRegister(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.CreateUser(user, ctrl.usersCollection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// HandleLogin authenticates a user and returns a JWT
func (ctrl *Controller) HandleLogin(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := data.GetUserByUsername(user.Username, ctrl.usersCollection)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !data.CheckPasswordHash(user.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateJWT(foundUser)
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

	if err := data.PromoteUser(id, ctrl.usersCollection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}
