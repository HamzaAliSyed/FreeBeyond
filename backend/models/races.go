package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Race struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	MovementSpeed      map[string]int     `bson:"movementspeed"`
	Rarity             string             `bson:"rarity"`
	Family             string             `bson:"family,omitempty"`
	Size               string             `bson:"size"`
	StatsBoost         map[string]int     `bson:"statsboost"`
	Languages          []string           `bson:"languages"`
	FlavourText        []FlavourText      `bson:"flavourtext"`
	SkillProficiencies []string           `bson:"skillproficiencies,omitempty"`
	AdvantageSkills    []string           `bson:"advantageskills,omitempty"`
	Attacks            []string           `bson:"attacks,omitempty"`
	Spells             map[string]int     `bson:"spells,omitempty"`
	Immunities         []string           `bson:"immunities,omitempty"`
	Resistances        []string           `bson:"resistances,omitempty"`
	PhysicalFeatures   map[string]string  `bson:"physicalfeatures"`
	SavingThrows       map[string]string  `bson:"savingthrows,omitempty"`
	Source             primitive.ObjectID `bson:"source"`
}
