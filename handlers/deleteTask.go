package handlers

import (
	"context"
	"net/http"
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/utils" 

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTask(c *gin.Context) {
	ctx := context.Background()
	idParam := c.Param("id")

	// Convert the task ID from string to ObjectID
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	// Get claims from the context (user info from the token)
	claims, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Type assertion to get user claims
	userClaims, ok := claims.(*utils.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
		return
	}

	// Convert user ID to ObjectID
	userID, err := primitive.ObjectIDFromHex(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// Build the filter to check both task ID and user ID
	filter := bson.M{
		"_id":    id,
		"userId": userID, // Ensure the task is owned by the user
	}

	// Perform the delete operation
	res, err := db.TaskCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	// If no task was deleted (not found or unauthorized)
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or unauthorized"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
