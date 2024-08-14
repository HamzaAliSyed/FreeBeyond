package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	UserID           primitive.ObjectID `bson:"user_id"`
	Name             string             `bson:"name"`
	PlayerName       string             `bson:"playername,omitempty"`
	ProficiencyBonus int                `bson:"proficiencybonus"`
	Inspiration      int                `bson:"inspiration"`
	ArmourClass      int                `bson:"armourclass"`
	LandSpeed        int                `bson:"landspeed"`
	FlyingSpeed      int                `bson:"flyingspeed"`
	SwimmingSpeed    int                `bson:"swimmingspeed"`
	ClimbingSpeed    int                `bson:"climbingspeed"`
	BurrowingSpeed   int                `bson:"burrowingspeed"`
	ExhaustionLevel  int                `bson:"exhaustionlevel"`
	MainAttributes   MainAttributes     `bson:"mainattributes"`
	Modifiers        Modifiers          `bson:"modifiers"`
	SavingThrow      []SavingThrow      `bson:"savingthrow"`
}
