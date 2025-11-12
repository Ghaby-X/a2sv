// Package controllers interacts with request and handle http related processing
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	services "github.com/Ghaby-X/task_manager/data"
	"github.com/Ghaby-X/task_manager/models"
	"github.com/gin-gonic/gin"
)

// HandleGetTask to handle get tasks
func HandleGetTask(c *gin.Context) {
	tasks := services.GetTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// HandleCreateTask create tasks
func HandleCreateTask(c *gin.Context) {
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
	err = services.CreateTask(newTask)
	if err != nil {
		println("failed to create task")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "task created successfully"})
}

// HandleGetTaskByID gets a task by id
func HandleGetTaskByID(c *gin.Context) {
	id := c.Param("id")

	num, err := strconv.Atoi(id)

	// check if id exist
	if len(id) < 1 || err != nil {
		fmt.Println("Could not parse id to number, id: ", id)
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to parse id"})
		return
	}

	task, err := services.GetTaskByID(num)
	if err != nil {
		fmt.Println("task not found, id: ", id)
		c.JSON(http.StatusOK, gin.H{"msg": "task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func HandleDeleteTask(c *gin.Context) {
	id := c.Param("id")

	num, err := strconv.Atoi(id)

	// check if id exist
	if len(id) < 1 || err != nil {
		fmt.Println("Could not parse id to number, id: ", id)
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to parse id"})
		return
	}

	err = services.DeleteTask(num)
	if err != nil {
		fmt.Println("failed to delete task")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "task deleted successfully"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)

	// check if id exist
	if len(id) < 1 || err != nil {
		fmt.Println("Could not parse id to number, id: ", id)
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to parse id"})
		return
	}

	// update task
	var newTask models.Task

	err = c.BindJSON(&newTask)
	if err != nil {
		fmt.Println("failed to parse task")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "failed to parse task json"})
		return
	}

	err = services.UpdateTask(num, newTask)
	if err != nil {
		fmt.Println("failed to update task")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "task updated successfully"})
}
