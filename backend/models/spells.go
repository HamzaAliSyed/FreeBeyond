package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SchoolOfMagic string

const (
	Abjuration    SchoolOfMagic = "Abjuration"
	Conjuration   SchoolOfMagic = "Conjuration"
	Divination    SchoolOfMagic = "Divination"
	Enchantment   SchoolOfMagic = "Enchantment"
	Evocation     SchoolOfMagic = "Evocation"
	Illusion      SchoolOfMagic = "Illusion"
	Necromancy    SchoolOfMagic = "Necromancy"
	Transmutation SchoolOfMagic = "Transmutation"
)

type Spells struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Level         int                `bson:"level"`
	CastingTime   string             `bson:"castingtime"`
	Duration      string             `bson:"duration"`
	School        SchoolOfMagic      `bson:"school"`
	Concentration bool               `bson:"concentration"`
	Range         string             `bson:"range"`
	Components    []string           `bson:"components"`
	FlavourText   string             `bson:"flavourtext"`
	Classes       string             `bson:"classes"`
	SubClasses    string             `bson:"subclasses"`
	Source        primitive.ObjectID `bson:"source"`
	SourceName    string             `bson:"sourcename"`
}

type AttackBasedRangeAOEAttack struct {
	Spells
	AOEShape      string            `bson:"aoeshape"`
	AOERadius     int               `bson:"aoeradius"`
	SaveAttribute string            `bson:"saveattribute"`
	Damage        map[string]string `bson:"damage"`
	SaveEffect    string            `bson:"saveeffect"`
}
