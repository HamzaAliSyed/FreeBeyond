package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`
	Name                      string              `bson:"name"`
	Levels                    []Level             `bson:"levels,omitempty"`
	SubClasses                []SubClass          `bson:"subclassess,omitempty"`
	UniqueFeatures            []string            `bson:"uniquefeatures,omitempty"`
	CanDoSpellCasting         bool                `bson:"candospellcasting"`
	Hitdie                    string              `bson:"hitdie"`
	ArmorProficiencies        []string            `bson:"armorproficiencies"`
	WeaponProficiencies       []string            `bson:"weaponproficiencies"`
	ToolProficiencies         []string            `bson:"toolproficiencies"`
	SavingThrowsProficiencies []string            `bson:"savingthrowsproficiencies"`
	SkillProficiencies        map[string][]string `bson:"skillproficiencies"`
	Source                    primitive.ObjectID  `bson:"source"`
}

type SubClass struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

type Level struct {
	Class                string              `bson:"class"`
	LevelRank            int                 `bson:"levelrank"`
	ProficiencyBonus     int                 `bson:"proficiencybonus"`
	Features             []ClassFeature      `bson:"features"`
	SpellSlots           map[int]int         `bson:"spellslots"`
	UniqueClassAbilities []map[string]string `bson:"uniqueclassabilities"`
}

type ClassFeature struct {
	Name                string               `bson:"name"`
	Type                string               `bson:"type"`
	Action              string               `bson:"action"`
	FeatureDie          string               `bson:"featuredie,omitempty"`
	TextInformation     string               `bson:"textinformation,omitempty"`
	ChargeBasedAbility  ChargeBasedAbilities `bson:"chargedbasedability,omitempty"`
	ResetInformation    string               `bson:"resetinformation,omitempty"`
	ModifierAbility     string               `bson:"modifierability,omitempty"`
	MappableAbility     map[string]int       `bson:"mappableability,omitempty"`
	TableAbility        map[string]string    `bson:"tableability,omitempty"`
	ConditionalFeatures []ClassFeature       `bson:"complexfeatures"`
}

type SubClassFeature struct {
	Name               string               `bson:"name"`
	TextInformation    FlavourText          `bson:"textinformation,omitempty"`
	ChargeBasedAbility ChargeBasedAbilities `bson:"chargedbasedability,omitempty"`
	ResetInformation   string               `bson:"resetinformation,omitempty"`
	ModifierAbility    string               `bson:"modifierability,omitempty"`
	MappableAbility    map[string]int       `bson:"mappableability,omitempty"`
	TableAbility       map[string]string    `bson:"tableability,omitempty"`
}
