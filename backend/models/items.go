package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemType string

const (
	Mundane   ItemType = "Mundane"
	Common    ItemType = "Common"
	Uncommon  ItemType = "Uncommon"
	Rare      ItemType = "Rare"
	VeryRare  ItemType = "Very Rare"
	Legendary ItemType = "Legendary"
	Artifact  ItemType = "Artifact"
)

type Items struct {
	ID                 primitive.ObjectID   `bson:"_id,omitempty"`
	Name               string               `bson:"name"`
	TypeTags           []string             `bson:"typetags"`
	ItemType           ItemType             `bson:"itemtype"`
	RequiresAttunement bool                 `bson:"requiresattunement"`
	Cost               string               `bson:"cost"`
	Weight             string               `bson:"weight"`
	FlavourText        []TextBasedAbility   `bson:"flavourtext,omitempty"`
	ChargeBasedAbility []ChargeBasedAbility `bson:"chargebasedability,omitempty"`
	ModifyAbility      []ModifierAbility    `bson:"modifyability,omitempty"`
	Spells             []Spells             `bson:"spells,omitempty"`
	Source             primitive.ObjectID   `bson:"source"`
}
