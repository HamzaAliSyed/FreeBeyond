package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Preparation string

const (
	InstantLearn Preparation = "InstantLearn"
	DailyChosen  Preparation = "DailyChosen"
)

type RuleType string

const (
	CoreRule RuleType = "CoreRule"
	Variant  RuleType = "Variant"
	Optional RuleType = "Optional"
)

type Levels struct {
	Level              int                  `bson:"level"`
	ProficiencyBonus   int                  `bson:"proficiencybonus"`
	FlavorAbilities    []TextBasedAbility   `bson:"flavourabilities,omitempty"`
	ChargeBasedAbility []ChargeBasedAbility `bson:"chargebasedability,omitempty"`
	ModifierAbility    []ModifierAbility    `bson:"modifierability,omitempty"`
	SpellCasting       SpellCasting         `bson:"spellcasting,omitempty"`
}

type TextBasedAbility struct {
	Title        string   `bson:"title"`
	FlavourText  string   `bson:"flavourtext"`
	Availability string   `bson:"availability"`
	RuleType     RuleType `bson:"ruletype"`
}

type ChargeBasedAbility struct {
	Title                    string   `bson:"title"`
	FlavourText              string   `bson:"flavourtext"`
	TotalCharges             int      `bson:"totalcharges"`
	NumberOfChargesRemaining int      `bson:"numberofchargesremaining"`
	Availability             string   `bson:"availability"`
	RuleType                 RuleType `bson:"ruletype"`
}

type ModifierAbility struct {
	Title              string   `bson:"title"`
	FlavourText        string   `bson:"flavourtext"`
	ModifierStatFamily string   `bson:"modifierstatfamily"`
	ModifierStat       string   `bson:"modifierstat"`
	ModifierValue      string   `bson:"modifiervalue"`
	Availability       string   `bson:"availability"`
	RuleType           RuleType `bson:"ruletype"`
}

type SpellCasting struct {
	Origin         primitive.ObjectID `bson:"origin"`
	Level          int                `bson:"level"`
	Preparation    Preparation        `bson:"preparation"`
	MajorAttribute string             `bson:"majorattribute"`
	SpellSlots     map[int]int        `bson:"spellslots"`
	Spells         []Spells           `bson:"spells"`
}

type SubClasses struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	ParentClass primitive.ObjectID `bson:"parentclass"`
	Levels      []Levels           `bson:"levels"`
	Source      primitive.ObjectID `bson:"source"`
}

type Class struct {
	ID                     primitive.ObjectID   `bson:"_id,omitempty"`
	Name                   string               `bson:"name"`
	HitDie                 string               `bson:"hitdie"`
	ArmorProficiency       []string             `bson:"armorproficiency"`
	WeaponProficiency      []string             `bson:"weaponproficiency"`
	ToolsProficiency       []string             `bson:"toolsproficiency"`
	SavingThrowProficiency []string             `bson:"savingthrowproficiency"`
	SkillsCanChoose        int                  `bson:"skillscanchoose"`
	SkillsChoiceList       []string             `bson:"skills"`
	ToolProficiencies      []primitive.ObjectID `bson:"toolproficiencies"`
	Levels                 []Levels             `bson:"levels"`
	SubClasses             []SubClasses         `bson:"subclasses"`
	SubClassChoices        []int                `bson:"subclasschoices"`
	Source                 primitive.ObjectID   `bson:"source"`
}
