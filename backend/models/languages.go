package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Languages struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Lore            string             `bson:"lore,omitempty"`
	Type            string             `bson:"type,omitempty"`
	TypicalSpeakers primitive.ObjectID `bson:"typicalspeakers,omitempty"`
	Script          string             `bson:"script,omitempty"`
	Source          primitive.ObjectID `bson:"source"`
}
