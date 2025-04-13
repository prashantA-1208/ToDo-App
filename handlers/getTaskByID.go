package handlers

import (
	"context"
	"net/http"
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTaskByID(c *gin.Context) {
	ctx := context.Background()
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var task models.Task
	err = db.TaskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}
