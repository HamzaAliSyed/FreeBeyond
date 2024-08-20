package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`
	Name                      string              `bson:"name"`
	CanDoSpellCasting         bool                `bson:"candospellcasting"`
	Hitdie                    string              `bson:"hitdie"`
	ArmorProficiencies        []string            `bson:"armorproficiencies"`
	WeaponProficiencies       []string            `bson:"weaponproficiencies"`
	ToolProficiencies         []string            `bson:"toolproficiencies"`
	SavingThrowsProficiencies []string            `bson:"savingthrowsproficiencies"`
	SkillProficiencies        map[string][]string `bson:"skillproficiencies"`
	Subclasses                []string            `bson:"subclasses"`
	ProficiencyBonusTable     map[int]int         `bson:"proficiencybonustable"`
	SpecialFeatureTable       []map[string]string `bson:"specialfeaturetable"`
	Source                    primitive.ObjectID  `bson:"source"`
}
