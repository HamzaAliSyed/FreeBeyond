package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item interface {
	Create() interface{}
	GetAllProperties() map[string]interface{}
}

type Weapon struct {
	ID primitive.ObjectID `bson:"_id"`
	BasicProperties
	rangemin int               `bson:"rangemin"`
	rangemax int               `bson:"rangemax"`
	damage   map[Damage]string `bson:"damage"`
}

func (weapon *Weapon) Create() interface{} {
	return weapon
}

type Rarity string

const (
	Mundane   Rarity = "Mundane"
	Common    Rarity = "Common"
	Uncommon  Rarity = "Uncommon"
	Rare      Rarity = "Rare"
	VeryRare  Rarity = "Very Rare"
	Legendary Rarity = "Legendary"
	Artifact  Rarity = "Artifact"
)

type Tier string

const (
	Insignificant Tier = "Insignificant"
	Major         Tier = "Major"
	Minor         Tier = "Minor"
)

type BasicProperties struct {
	name               string   `bson:"name"`
	typetags           []string `bson:"typetags"`
	rarity             Rarity   `bson:"rarity"`
	tier               Tier     `bson:"tier"`
	requiresAttunement bool     `bson:"requiresattunement"`
	description        string   `bson:"description"`
	cost               string   `bson:"cost"`
	weight             string   `bson:"weight"`
	source             string   `bson:"source"`
}

func isValidRarity(r Rarity) bool {
	validRarities := []Rarity{Mundane, Common, Uncommon, Rare, VeryRare, Legendary, Artifact}
	for _, v := range validRarities {
		if r == v {
			return true
		}
	}
	return false
}

func isValidTier(t Tier) bool {
	validTiers := []Tier{Insignificant, Major, Minor}
	for _, v := range validTiers {
		if t == v {
			return true
		}
	}

	return false
}

func isValidDamage(d Damage) bool {
	validDamage := []Damage{Acid, Bludgeoning, Cold, Fire, Force, Lightning, Necrotic, Piercing, Poison, Psychic, Radiant, Slashing, Thunder}
	for _, damage := range validDamage {
		if damage == d {
			return true
		}
	}

	return false
}

func CreateNewWeapon(weaponName, weaponDescription, weaponCost, weaponWeight, weaponSource string, weaponTypeTags []string, weaponRarity Rarity, weaponTier Tier, needsAtunement bool, weaponRangeMin, weaponRangeMax int, weaponDamage map[Damage]string) (*Weapon, error) {

	weapon := &Weapon{}

	for damageType := range weaponDamage {
		if !isValidDamage(damageType) {
			return nil, fmt.Errorf("invalid damage type: %v", damageType)
		}
	}

	if !isValidRarity(weaponRarity) {
		return nil, errors.New("invalid rarity")
	}

	if !isValidTier(weaponTier) {
		return nil, errors.New("invalid tier")
	}

	weapon.name = weaponName
	weapon.typetags = weaponTypeTags
	weapon.rarity = weaponRarity
	weapon.tier = weaponTier
	weapon.requiresAttunement = needsAtunement
	weapon.description = weaponDescription
	weapon.cost = weaponCost
	weapon.weight = weaponWeight
	weapon.source = weaponSource
	weapon.rangemin = weaponRangeMin
	weapon.rangemax = weaponRangeMax
	weapon.damage = weaponDamage

	return weapon, nil
}

func (w *Weapon) GetAllProperties() map[string]interface{} {
	return map[string]interface{}{
		"ID":                 w.ID,
		"name":               w.name,
		"typetags":           w.typetags,
		"rarity":             w.rarity,
		"tier":               w.tier,
		"requiresAttunement": w.requiresAttunement,
		"description":        w.description,
		"cost":               w.cost,
		"weight":             w.weight,
		"source":             w.source,
		"rangemin":           w.rangemin,
		"rangemax":           w.rangemax,
		"damage":             w.damage,
	}
}
