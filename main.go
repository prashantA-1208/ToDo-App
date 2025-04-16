package main

import (
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/handlers"
	"prashantA-1208/ToDo-App.git/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()

	router := gin.Default()

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/tasks", handlers.GetTasks)
		authorized.GET("/tasks/:id", handlers.GetTaskByID)
		authorized.POST("/tasks", handlers.CreateTask)
		authorized.PUT("/tasks/:id", handlers.UpdateTask)
		authorized.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	router.Run(":8080")

}
