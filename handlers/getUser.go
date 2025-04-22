package handlers

import (
	"context"
	"net/http"
	"prashantA-1208/ToDo-App.git/db"
	"prashantA-1208/ToDo-App.git/utils" // Ensure this import is correct
	"prashantA-1208/ToDo-App.git/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *gin.Context) {
	ctx := context.Background()

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
		"_id": userID,
	}

	var user models.User
	err = db.UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found or unauthorized"})
		return
	}

	c.JSON(http.StatusOK, user)
}
