package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FlavourText struct {
	Parent  string
	Heading string
	Body    string
}

type CharacterStatsAndAbilitiesModifier struct {
	Parent    string
	Attribute string
	Value     int
}

type ChargeBasedAbilities struct {
	Parent             string
	Name               string
	MaxNumberOfCharges int
	NumberOfCharges    int
}

type AnAttack struct {
	Parent           string
	Name             string
	Type             string
	Range            int
	RangeMax         int
	AttributeUsed    string
	AttributeValue   int
	IsProficient     bool
	ProficiencyBonus int
	Damage           map[string]string
	OtherBonus       string
	OtherBonusValue  int
}

type SpellAttack struct {
	Parent         string
	NameOfSpell    string
	Level          int
	Classes        []primitive.ObjectID
	TypeOfSpell    string
	SourceOfSpell  string
	Range          int
	AreaOfEffect   int
	NeedsDC        bool
	DCValue        int
	NeedsAttack    bool
	AttackValue    int
	Damage         map[string]string
	SpellReference primitive.ObjectID
}
