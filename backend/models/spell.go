package models

type School string

const (
	Abjuration    School = "Abjuration"
	Conjuration   School = "Conjuration"
	Divination    School = "Divination"
	Enchantment   School = "Enchantment"
	Evocation     School = "Evocation"
	Illusion      School = "Illusion"
	Necromancy    School = "Necromancy"
	Transmutation School = "Transmutation"
)

type TypeOfSpell string

const (
	Generic      TypeOfSpell = "Generic"
	SaveBased    TypeOfSpell = "Save Based"
	RollToAttack TypeOfSpell = "Roll To Attack"
	Modifier     TypeOfSpell = "Modifier"
)

type Damage map[string]string

type Spell struct {
	Name                  string            `bson:"name"`
	FlavourText           string            `bson:"flavourtext"`
	Level                 int               `bson:"level"`
	ActionTime            string            `bson:"actiontime"`
	CastingTime           string            `bson:"castingtime"`
	Components            []string          `bson:"components"`
	School                School            `bson:"school"`
	RequiresConcentration bool              `bson:"requiresconcentration"`
	TypeOfSpell           TypeOfSpell       `bson:"typeofspell"`
	Range                 string            `bson:"range"`
	AccessTo              map[string]string `bson:"accessto"`
	Source                string            `bson:"source"`
	Shape                 string            `bson:"shape,omitempty"`
	ShapeSize             string            `bson:"shapesize,omitempty"`
	Damage                Damage            `bson:"damage,omitempty"`
	SaveStat              string            `bson:"savestat,omitempty"`
	SaveEffect            string            `bson:"saveeffect,omitempty"`
	StatsModified         []string          `bson:"statsmodified,omitempty"`
	NumberofCreature      string            `bson:"numberofcreature,omitempty"`
}
