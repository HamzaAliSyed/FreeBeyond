package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	UserID           primitive.ObjectID `bson:"user_id"`
	Name             string             `bson:"name"`
	PlayerName       string             `bson:"playername,omitempty"`
	ProficiencyBonus int                `bson:"proficiencybonus"`
	Inspiration      int                `bson:"inspiration"`
	Race             string             `bson:"race"`
	ClassAndLevel    map[string]int     `bson:"classandlevel"`
	HitDie           string             `bson:"hitdie"`
	NumberOfHitDie   int                `bson:"numberofhitdie"`
	HealthPoints     int                `bson:"healthpoints"`
	TempPoints       int                `bson:"temppoints"`
	ArmourClass      int                `bson:"armourclass"`
	LandSpeed        int                `bson:"landspeed"`
	FlyingSpeed      int                `bson:"flyingspeed"`
	SwimmingSpeed    int                `bson:"swimmingspeed"`
	ClimbingSpeed    int                `bson:"climbingspeed"`
	BurrowingSpeed   int                `bson:"burrowingspeed"`
	Initiative       int                `bson:"initiative"`
	Background       string             `bson:"background"`
	Appearance       map[string]string  `bson:"appearance,omitempty"`
	ExhaustionLevel  int                `bson:"exhaustionlevel"`
	MainAttributes   MainAttributes     `bson:"mainattributes"`
	Modifiers        Modifiers          `bson:"modifiers"`
	SavingThrow      []SavingThrow      `bson:"savingthrow"`
	Skills           Skills             `bson:"skills"`
	CharacterMotives CharacterMotives   `bson:"charactermotives"`
	Feats            []Feats            `bson:"feats"`
	MaxCarryWeight   int                `bson:"maxcarryweight"`
	CarryWeight      int                `bson:"carryweight"`
	Inventory        []Items            `bson:"inventory"`
	StatusAfflicted  []string           `bson:"statusafflicted"`
	Languages        []string           `bson:"languages"`
	Attacks          []AnAttack         `bson:"attack"`
	CanSpellCast     bool               `bson:"canspellcast"`
	SpellsAvailable  []SpellAttack      `bson:"spellsavailable"`
	Features         []string           `bson:"features"`
	Currency         map[string]int     `bson:"currency"`
	Proficiencies    map[string]string  `bson:"proficiencies"`
	OtherAbilities   []GenericAbility   `bson:"otherabilities,omitempty"`
}
