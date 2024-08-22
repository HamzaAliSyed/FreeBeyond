package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type school string

const (
	Abjuration    school = "Abjuration"
	Conjuration   school = "Conjuration"
	Divination    school = "Divination"
	Enchantment   school = "Enchantment"
	Evocation     school = "Evocation"
	Illusion      school = "Illusion"
	Necromancy    school = "Necromancy"
	Transmutation school = "Transmutation"
)

type Spells struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Name                   string             `bson:"name"`
	Level                  int                `bson:"level"`
	Time                   string             `bson:"time"`
	Action                 string             `bson:"action"`
	School                 school             `bson:"school"`
	Concentration          bool               `bson:"concentration"`
	Range                  string             `bson:"range"`
	Source                 string             `bson:"source"`
	SourceID               primitive.ObjectID `bson:"sourceid"`
	Components             []string           `bson:"components"`
	Description            string             `bson:"description"`
	Duration               string             `bson:"duration,omitempty"`
	TypeOfSpell            string             `bson:"typeofspell,omitempty"` //Attack,Heal,Other
	AttackTypeOfSpell      string             `bson:"AttackTypeOfSpell,omitempty"`
	DamageReductionOfSpell string             `bson:"damagereductionofspell,omitempty"` //if miss then what happens
	DieValue               string             `bson:"dievalue,omitempty"`
	Classes                []string           `bson:"classes"`
	SubClasses             []string           `bson:"subclasses"`
	Effect                 string             `bson:"effect,omitempty"`
}
