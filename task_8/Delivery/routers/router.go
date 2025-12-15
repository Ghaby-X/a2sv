
package routers

import (
	"github.com/Ghaby-X/task_manager/Delivery/controllers"
	"github.com/Ghaby-X/task_manager/Infrastructure"
	"github.com/gin-gonic/gin"
)

func GetTaskRouter(ctrl *controllers.Controller, jwtService Infrastructure.JWTService) *gin.Engine {
	r := gin.Default()

	// Public routes for user registration and login
	r.POST("/register", ctrl.HandleRegister)
	r.POST("/login", ctrl.HandleLogin)

	// Group for authenticated routes
	authenticated := r.Group("/")
	authenticated.Use(Infrastructure.AuthMiddleware(jwtService))
	{
		// Task routes (accessible by all authenticated users)
		authenticated.GET("/tasks", ctrl.HandleGetTask)
		authenticated.GET("/tasks/:id", ctrl.HandleGetTaskByID)

		// Admin-only task routes
		adminTasks := authenticated.Group("/tasks")
		adminTasks.Use(Infrastructure.AuthorizeRole("admin"))
		{
			adminTasks.POST("/", ctrl.HandleCreateTask)
			adminTasks.PUT("/:id", ctrl.UpdateTask)
			adminTasks.DELETE("/:id", ctrl.HandleDeleteTask)
		}

		// Admin-only user promotion route
		adminUsers := authenticated.Group("/users")
		adminUsers.Use(Infrastructure.AuthorizeRole("admin"))
		{
			adminUsers.POST("/promote/:id", ctrl.HandlePromoteUser)
		}
	}

	return r
}
