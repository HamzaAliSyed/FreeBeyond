package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	FirstName  string               `bson:"first_name"`
	LastName   string               `bson:"last_name"`
	Username   string               `bson:"username"`
	Email      string               `bson:"email"`
	Password   string               `bson:"password"`
	Role       string               `bson:"role"`
	Characters []primitive.ObjectID `bson:"characters"`
}
