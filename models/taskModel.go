package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string
	Completed bool
}
