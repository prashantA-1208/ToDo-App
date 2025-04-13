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

func UpdateTask(c *gin.Context) {
	ctx := context.Background()
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updated models.Task
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":     updated.Title,
			"completed": updated.Completed,
		},
	}

	_, err = db.TaskCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}
