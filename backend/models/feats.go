package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feats struct {
	ID                     primitive.ObjectID                   `bson:"_id,omitempty"`
	FeatName               string                               `bson:"featname"`
	FeatDescription        string                               `bson:"featdescription"`
	Prerequisite           string                               `bson:"prequisite,omitempty"`
	TextFeature            []FlavourText                        `bson:"textfeature,omitempty"`
	CharacterModifications []CharacterStatsAndAbilitiesModifier `bson:"charactermodification,omitempty"`
	ChargeAbilities        []ChargeBasedAbilities               `bson:"chargeabilities,omitempty"`
	Attacks                []AnAttack                           `bson:"attack,omitempty"`
	Spells                 []SpellAttack                        `bson:"spells,omitempty"`
	Source                 primitive.ObjectID                   `bson:"objectID"`
}
