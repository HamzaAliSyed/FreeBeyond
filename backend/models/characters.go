package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Damage string

const (
	Slashing    Damage = "Slashing"
	Piercing    Damage = "Piercing"
	Bludgeoning Damage = "Bludgeoning"
	Fire        Damage = "Fire"
	Cold        Damage = "Cold"
	Lightning   Damage = "Lightning"
	Thunder     Damage = "Thunder"
	Acid        Damage = "Acid"
	Poison      Damage = "Poison"
	Force       Damage = "Force"
	Psychic     Damage = "Psychic"
	Necrotic    Damage = "Necrotic"
	Radiant     Damage = "Radiant"
)

type Character struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	ClassAndLevel     map[string]int     `bson:"classandlevel,omitempty"`
	Allignment        string             `bson:"allignment"`
	ProficiencyBonus  int                `bson:"proficiencybonus"`
	PersonalityTrait  string             `bson:"personalitytrait"`
	Ideals            string             `bson:"ideals"`
	Bonds             string             `bson:"bonds"`
	Flaws             string             `bson:"flaws"`
	PhysicalFeatures  map[string]string  `bson:"physicalfeatures"`
	AbilityScores     []AbilityScore     `bson:"abilityscores"`
	SavingThrows      []SavingThrow      `bson:"savingthrows"`
	CharacterImage    primitive.Binary   `bson:"characterimage"`
	Movements         Movement           `bson:"movement"`
	Health            Health             `bson:"health"`
	Attacks           []Attack           `bson:"attacks"`
	Race              primitive.ObjectID `bson:"race"`
	DamageResistance  []Damage           `bson:"damageresistance"`
	DamageImmunities  []Damage           `bson:"damageimmunities"`
	WeaponProficiency []string           `bson:"weaponproficiency"`
	ArmorProficiency  []string           `bson:"armorproficiency"`
	ToolsProficiency  []string           `bson:"toolsproficiency"`
	Senses            []Sense            `bson:"senses"`
	Languages         []string           `bson:"langugaes"`
}

type AbilityScore struct {
	AbilityName     string `bson:"abilityname"`
	AbilityScore    int    `bson:"abilityscore"`
	AbilityModifier int    `bson:"abilitymodifier"`
}

type Movement struct {
	Initiative  int `bson:"initiative"`
	LandSpeed   int `bson:"landspeed"`
	SwimSpeed   int `bson:"swimspeed"`
	FlySpeed    int `bson:"flyspeed"`
	ClimbSpeed  int `bson:"climbspeed"`
	BurrowSpeed int `bson:"burrowspeed"`
}

type Health struct {
	MaxHealth     int `bson:"maxhealth"`
	CurrentHealth int `bson:"currenthealth"`
	TempHealth    int `bson:"temphealth"`
}

type SavingThrow struct {
	Ability               string  `bson:"ability"`
	NumberOfProficiencies float64 `bson:"numberofproficiencies"`
	HasAdvantage          bool    `bson:"hasadvantage"`
	HasDisadvantage       bool    `bson:"hasdisadvantage"`
	Value                 int     `bson:"value"`
}

type Attack struct {
	AttackName string            `bson:"attackname"`
	Attribute  string            `bson:"attribute"`
	Damage     map[string]string `bson:"damage"`
	Range      int               `bson:"range"`
}

type Sense struct {
	SenseName  string `bson:"sensename"`
	SenseRange int    `bson:"senserange"`
}
