package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Source struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Type        string             `bson:"type"`
	PublishDate time.Time          `bson:"publishdate"`
}

func (source *Source) FormatPublishDate() string {
	return source.PublishDate.Format("January-02-2006")
}
