package models

type CharacterMotives struct {
	PersonalityTraits string `bson:"personalitytraits"`
	Ideals            string `bson:"ideals"`
	Bonds             string `bson:"bonds"`
	Flaws             string `bson:"flaws"`
}
