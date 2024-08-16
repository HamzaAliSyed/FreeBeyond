package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Skill struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Name                     string             `bson:"name"`
	AssociatedAttribute      string             `bson:"associatedattribute"`
	AssociatedAttributeValue int                `bson:"associatedattributevalue"`
	NumberOfProficiencies    float64            `bson:"numberofproficiencies"`
	ProficiencyBonus         int                `bson:"proficiencybonus"`
	AdditionalBoost          string             `bson:"additionalboost"`
	AdditionalBoostValue     int                `bson:"additionalboostvalue"`
	FinalSkillValue          int                `bson:"skillvalue"`
}

type Skills struct {
	SkillList []Skill `bson:"skilllist"`
}
