package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Source struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Type        string             `bson:"type"`
	PublishDate string             `bson:"publishdate"`
}
