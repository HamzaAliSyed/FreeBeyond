package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Items struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Rarity string             `bson:"category"`
	Tags   []string           `bson:"tags"`
	Cost   int                `bson:"cost,omitempty"`
	Weight int                `bson:"weight,omitempty"`
	Source primitive.ObjectID `bson:"source"`
}
