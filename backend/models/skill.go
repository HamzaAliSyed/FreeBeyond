package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Skill struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Name                string             `bson:"name"`
	AssociatedAttribute string             `bson:"associatedattribute"`
	ProficiencyBonus    int                `bson:"proficiencybonus"`
	SkillValue          int                `bson:"skillvalue"`
}

type Skills struct {
	SkillList []Skill `bson:"skilllist"`
}
