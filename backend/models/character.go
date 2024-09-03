package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HitDie string

type ClassData struct {
	ClassesAndLevels map[string]int
	HitDices         map[HitDie]int
}

type AbilityScore struct {
	AbilityName string `bson:"abilityname"`
	Score       int    `bson:"score"`
	Modifier    int    `bson:"modifier"`
}

type SavingThrow struct {
	AttributeName      string  `bson:"attributename"`
	Modifer            int     `bson:"modiifer"`
	AdditionalBonus    int     `bson:"additionalbonus"`
	NumberOProficiency float64 `bson:"numberofproficiency"`
	Value              int     `bson:"value"`
	HasAdvantage       bool    `bson:"hasadvantage"`
	HasDisadvantage    bool    `bson:"hasdisadvantage"`
}

type Skill struct {
	AttributeName      string  `bson:"attributename"`
	Modifer            int     `bson:"modiifer"`
	AdditionalBonus    int     `bson:"additionalbonus"`
	NumberOProficiency float64 `bson:"numberofproficiency"`
	Value              int     `bson:"value"`
	HasAdvantage       bool    `bson:"hasadvantage"`
	HasDisadvantage    bool    `bson:"hasdisadvantage"`
}

type Character struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`
	Class            ClassData          `bson:"class"`
	ProficiencyBonus int                `bson:"proficiencybonus"`
	AbilityScores    []AbilityScore     `bson:"abilityscores"`
	SavingThrows     []SavingThrow      `bson:"savingthrows"`
	Skills           []Skill            `bson:"skills"`
}
