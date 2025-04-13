package handlers

import (
	"context"
	"net/http"
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(c *gin.Context) {
	ctx := context.Background()

	cursor, err := db.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
