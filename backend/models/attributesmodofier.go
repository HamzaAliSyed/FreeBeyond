package models

type Modifiers struct {
	StrengthModifier     int `bson:"strengthmodifier"`
	DexterityModifier    int `bson:"dexteritymodifier"`
	ConstitutionModifier int `bson:"constitutionmodifier"`
	IntelligenceModifier int `bson:"intelligencemodifier"`
	WisdomModifier       int `bson:"wisdommodifier"`
	CharismaModifier     int `bson:"charismamodifier"`
}
