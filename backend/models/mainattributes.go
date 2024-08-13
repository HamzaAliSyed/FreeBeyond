package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MainAttributes struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	StrengthScore     int                `bson:"strengthscore"`
	DexterityScore    int                `bson:"dexterityscore"`
	ConstitutionScore int                `bson:"constitutionscore"`
	IntelligenceScore int                `bson:"intelligencescore"`
	WisdomScore       int                `bson:"wisdomscore"`
	CharismaScore     int                `bson:"charismascore"`
}
