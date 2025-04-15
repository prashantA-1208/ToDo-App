package handlers

import (
	"context"
	"net/http"
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {
	ctx := context.Background()
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task.ID = primitive.NewObjectID()

	_, err := db.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	userID, exists := c.Get("userId")
if !exists {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	return
}

uid, _ := primitive.ObjectIDFromHex(userID.(string))
task.UserID = uid

	c.JSON(http.StatusCreated, task)
}
