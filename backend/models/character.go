package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HitDie string

const (
	d6  HitDie = "d6"
	d8  HitDie = "d8"
	d10 HitDie = "d10"
	d12 HitDie = "d12"
)

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

type Initiative struct {
	DexterityModifier int `bson:"dexteritymodifier"`
	AdditionalBonus   int `bson:"additionalbonus"`
	Value             int `bson:"value"`
}

type Movement struct {
	LandSpeed   int `bson:"landspeed"`
	SwimSpeed   int `bson:"swimspeed"`
	ClimbSpeed  int `bson:"climbspeed"`
	BurrowSpeed int `bson:"burrowspeed"`
	FlySpeed    int `bson:"flyspeed"`
}

type Sense struct {
	SenseName  string `bson:"sensename"`
	SenseRange int    `bson:"senserange"`
}

type EquippedItems struct {
	Head    string `bson:"head"`
	Torso   string `bson:"torso"`
	Legs    string `bson:"legs"`
	OffHand string `bson:"offhand"`
	Hand    string `bson:"hand"`
	Wrist   string `bson:"wrist"`
	SideArm string `bson:"sidearm"`
	Neck    string `bson:"neck"`
}

type Attunement struct {
	MaxAttunement       int      `bson:"maxattunement"`
	NumberOfAttunements int      `bson:"numberofattunement"`
	AttunedItems        []string `bson:"attuneditems"`
}

type Character struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	Name                 string             `bson:"name"`
	Class                ClassData          `bson:"class"`
	ProficiencyBonus     int                `bson:"proficiencybonus"`
	AbilityScores        []AbilityScore     `bson:"abilityscores"`
	SavingThrows         []SavingThrow      `bson:"savingthrows"`
	Skills               []Skill            `bson:"skills"`
	Initiative           Initiative         `bson:"Initiative"`
	PassiveInvestigation int                `bson:"passiveinvestigation"`
	PassivePerception    int                `bson:"passiveperception"`
	PassiveInsight       int                `bson:"passiveinsight"`
	Movement             Movement           `bson:"movement"`
	Senses               []Sense            `bson:"senses"`
	EquippedItems        EquippedItems      `bson:"equippedItems"`
	Attunement           Attunement         `bson:"attunement"`
	Inventory            []Items            `bson:"inventory"`
	WeaponProficiencies  []string           `bson:"weaponproficiencies"`
	ArmorProficiencies   []string           `bson:"armorproficiencies"`
	ToolProficiencies    []string           `bson:"toolproficiencies"`
	Action               []string           `bson:"active"`
	Passive              []string           `bson:"passive"`
	Reaction             []string           `bson:"reaction"`
	Bonus                []string           `bson:"bonus"`
}
