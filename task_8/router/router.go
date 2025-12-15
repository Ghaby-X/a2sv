package router

import (
	"github.com/Ghaby-X/task_manager/controllers"
	"github.com/Ghaby-X/task_manager/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo" // New import
)

func GetTaskRouter(tasksCollection, usersCollection *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Create a new controller instance
	ctrl := controllers.NewController(tasksCollection, usersCollection)

	// Public routes for user registration and login
	r.POST("/register", ctrl.HandleRegister)
	r.POST("/login", ctrl.HandleLogin)

	// Group for authenticated routes
	authenticated := r.Group("/")
	authenticated.Use(middleware.AuthMiddleware())
	{
		// Task routes (accessible by all authenticated users)
		authenticated.GET("/tasks", ctrl.HandleGetTask)
		authenticated.GET("/tasks/:id", ctrl.HandleGetTaskByID)

		// Admin-only task routes
		adminTasks := authenticated.Group("/tasks")
		adminTasks.Use(middleware.AuthorizeRole("admin"))
		{
			adminTasks.POST("/", ctrl.HandleCreateTask)
			adminTasks.PUT("/:id", ctrl.UpdateTask)
			adminTasks.DELETE("/:id", ctrl.HandleDeleteTask)
		}

		// Admin-only user promotion route
		adminUsers := authenticated.Group("/users")
		adminUsers.Use(middleware.AuthorizeRole("admin"))
		{
			adminUsers.POST("/promote/:id", ctrl.HandlePromoteUser)
		}
	}

	return r
}