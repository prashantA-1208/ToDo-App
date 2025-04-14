package db

import (
	"context"
	"log"
	"time"

	"prashantA-1208/ToDo-App.git/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.GetMongoURI())

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	TaskCollection = Client.Database("ToDo-App").Collection("tasks")
	UserCollection = Client.Database("yourdbname").Collection("users")
}
