// Package router directs rounter to appropriate handler
package router

import (
	"github.com/Ghaby-X/task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func GetTaskRouter() *gin.Engine {
	r := gin.Default()

	// register routes
	r.GET("/tasks", controllers.HandleGetTask)
	r.GET("/tasks/:id", controllers.HandleGetTaskByID)

	r.POST("/tasks", controllers.HandleCreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)

	r.DELETE("/tasks/:id", controllers.HandleDeleteTask)

	return r
}
