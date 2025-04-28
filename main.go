package main

import (
	"time"
	"log"
	"os"

	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/handlers"
	"prashantA-1208/ToDo-App.git/middleware"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	db.Connect()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
		authorized.GET("/user", handlers.GetUser)
	}

	//router.Run(":8080")
	router.Run(":" + port)

}
