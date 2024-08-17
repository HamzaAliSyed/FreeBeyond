package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Species struct {
	ID                  primitive.ObjectID         `bson:"_id,omitempty"`
	Name                string                     `bson:"name"`
	Size                string                     `bson:"size"`
	LandSpeed           int                        `bson:"landspeed"`
	FlySpeed            int                        `bson:"flyspeed"`
	SwimmingSpeed       int                        `bson:"swimmingspeed"`
	BurrowSpeed         int                        `bson:"burrowingspeed"`
	ClimbingSpeed       int                        `bson:"climbingspeed"`
	AbilityImprovements map[string]int             `bson:"abilityscoreimprovements"`
	PhyicalAppearance   map[string]string          `bson:"physicalappearance"`
	SavingThrow         map[string]string          `bson:"savingthrows,omitempty"`
	Spells              map[primitive.ObjectID]int `bson:"spell,omitempty"`
	Proficiencies       map[string]int             `bson:"procificiencies,omitempty"`
	Resistances         []string                   `bson:"resistances,omitempty"`
	Immunities          []string                   `bson:"immunities,omitempty"`
	Attacks             []AnAttack                 `bson:"attacks,omitempty"`
}
