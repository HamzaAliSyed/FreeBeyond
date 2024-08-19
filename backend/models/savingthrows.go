package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SavingThrow struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Attribute             string             `bson:"attribute"`
	AttributeModifier     int                `bson:"attributemodifier"`
	NumberOfProficiencies int                `bson:"numberofproficiencies"`
	SavingThrowValue      int                `bson:"savingthrowvalue"`
	HasAdvantage          bool               `bson:"hasadvantage"`
}
