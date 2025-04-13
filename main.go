package main

import (
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()

	router := gin.Default()

	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTaskByID)
	router.POST("/tasks", handlers.CreateTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)

	router.Run(":8080")

}
