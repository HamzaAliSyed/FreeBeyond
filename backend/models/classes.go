package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Preparation string

const (
	InstantLearn Preparation = "InstantLearn"
	DailyChosen  Preparation = "DailyChosen"
)

type RuleType string

const (
	CoreRule RuleType = "Core Rule"
	Variant  RuleType = "Variant"
	Optional RuleType = "Optional"
)

type Recovery string

const (
	ShortRest Recovery = "Short Rest"
	LongRest  Recovery = "Long Rest"
)

type ChoiceOptions struct {
	List         int      `bson:"list"`
	ToChooseFrom []string `bson:"tochoosefrom"`
}

type Levels struct {
	Level              int                      `bson:"level"`
	ProficiencyBonus   int                      `bson:"proficiencybonus"`
	FlavorAbilities    []TextBasedAbility       `bson:"flavourabilities,omitempty"`
	ChargeBasedAbility []ChargeBasedAbility     `bson:"chargebasedability,omitempty"`
	ModifierAbility    []ModifierAbility        `bson:"modifierability,omitempty"`
	SpellCasting       SpellCasting             `bson:"spellcasting,omitempty"`
	Choices            map[string]ChoiceOptions `bson:"choices"`
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
	Recovery                 Recovery `bson:"recovery"`
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
	ArmorProficiency       []string             `bson:"armorproficiency,omitempty"`
	WeaponProficiency      []string             `bson:"weaponproficiency,omitempty"`
	ToolsProficiency       []string             `bson:"toolsproficiency,omitempty"`
	SavingThrowProficiency []string             `bson:"savingthrowproficiency,omitempty"`
	SkillsCanChoose        int                  `bson:"skillscanchoose,omitempty"`
	SkillsChoiceList       []string             `bson:"skills,omitempty"`
	ToolProficiencies      []primitive.ObjectID `bson:"toolproficiencies,omitempty"`
	Levels                 []Levels             `bson:"levels,omitempty"`
	SubClasses             []primitive.ObjectID `bson:"subclasses,omitempty"`
	SubClassChoices        []int                `bson:"subclasschoices,omitempty"`
	Source                 primitive.ObjectID   `bson:"source"`
}
