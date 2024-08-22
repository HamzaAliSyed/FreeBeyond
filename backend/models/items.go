package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GenericItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description []Description      `bson:"description"`
}

func (item GenericItem) String() string {
	return item.Name
}
