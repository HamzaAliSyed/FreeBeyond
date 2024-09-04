package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RuleType string
type FeatureType string
type Recharge string

const (
	CoreRule     RuleType = "Core"
	VariantRule  RuleType = "Variant"
	OptionalRule RuleType = "Optional"
)

const (
	PassiveFeature  FeatureType = "Passive"
	ActionFeature   FeatureType = "Action"
	BonusFeature    FeatureType = "Bonus"
	ReactionFeature FeatureType = "Reaction"
	FreeFeature     FeatureType = "Free"
)

const (
	AlwaysOn  Recharge = "AlwaysOn"
	ShortRest Recharge = "Short Rest"
	LongRest  Recharge = "Long Rest"
)

type ClassFeature struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	RuleType    RuleType           `bson:"ruletype"`
	FeatureType FeatureType        `bson:"featuretype"`
}

type StatsModifierFeature struct {
	ClassFeature       `bson:",inline"`
	AugmentationType   string   `bson:"augmentationtype"`
	AugmentationStats  []string `bson:"augmentationstats"`
	AugmentationValues []string `bson:"augmentationvalues"`
}
