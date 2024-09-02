package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Race struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	AbilityScores map[string]int     `bson:"abilityscore"`
	Size          string             `bson:"size"`
	Speed         map[string]int     `bson:"speed"`
	CreatureType  string             `bson:"creaturetype"`
	FlavourText   []TextBasedAbility `bson:"flavourtext"`
	Spells        map[string]int     `bson:"spells,omitempty"`
	Attacks       []Attack           `bson:"attacks,omitempty"`
	OtherBoost    map[string]string  `bson:"otherboost,omitempty"`
	AgeRange      []int              `bson:"agerange"`
	Languages     []string           `bson:"languages"`
	Image         primitive.Binary   `bson:"Image"`
	Source        primitive.ObjectID `bson:"source"`
}
