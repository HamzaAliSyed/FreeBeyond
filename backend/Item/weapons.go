package item

import (
	"errors"
	"fmt"
)

type Damage string

const (
	Acid        Damage = "Acid"
	Bludgeoning Damage = "Bludgeoning"
	Cold        Damage = "Cold"
	Fire        Damage = "Fire"
	Force       Damage = "Force"
	Lightning   Damage = "Lightning"
	Necrotic    Damage = "Necrotic"
	Piercing    Damage = "Piercing"
	Poison      Damage = "Poison"
	Psychic     Damage = "Psychic"
	Radiant     Damage = "Radiant"
	Slashing    Damage = "Slashing"
	Thunder     Damage = "Thunder"
)

func IsValidDamageType(d Damage) bool {
	switch d {
	case Acid, Bludgeoning, Cold, Fire, Force, Lightning, Necrotic, Piercing, Poison, Psychic, Radiant, Slashing, Thunder:
		return true
	}
	return false
}

type Weapon struct {
	BasicParameter
	weaponProperty            []string          `bson:"weaponproperty"`
	weaponPropertyDescription []string          `bson:"weaponpropertydescription"`
	damage                    map[Damage]string `bson:"damage"`
	rangemax                  int               `bson:"rangemax"`
	rangemin                  int               `bson:"rangemin"`
}

type WeaponParams struct {
	Name                      string
	Rarity                    Rarity
	TypeTags                  []string
	Description               string `bson:"description"`
	Cost                      string
	Weight                    string
	Source                    string
	DamageType                []Damage
	DamageAmount              []string
	WeaponProperty            []string
	WeaponPropertyDescription []string
	RangeMin                  int
	RangeMax                  int
}

func (w *Weapon) Create(params CreateParams) error {
	wp, ok := params.(WeaponParams)
	if !ok {
		return errors.New("invalid params for weapon")
	}
	w.name = wp.Name
	w.rarity = wp.Rarity
	w.damage = make(map[Damage]string)
	for i, dt := range wp.DamageType {
		w.damage[dt] = wp.DamageAmount[i]
	}
	w.typeTags = wp.TypeTags
	w.description = wp.Description
	w.cost = wp.Cost
	w.source = wp.Source
	w.weight = wp.Weight
	w.weaponProperty = wp.WeaponProperty
	w.weaponPropertyDescription = wp.WeaponPropertyDescription
	w.rangemin = wp.RangeMin
	w.rangemax = wp.RangeMax
	return nil
}

func (weapon *Weapon) Print() {
	fmt.Printf("Weapon Name: %v\n", weapon.name)
	for _, tag := range weapon.typeTags {
		fmt.Printf("Weapon Tags: %v\n", tag)
	}
	fmt.Printf("Weapon Description: %v\n", weapon.description)
	fmt.Printf("Rarity: %v\n", weapon.rarity)
	for _, weaponProperty := range weapon.weaponProperty {
		fmt.Printf("Weapon Property: %v", weaponProperty)
	}
	for _, weaponPropertDescription := range weapon.weaponPropertyDescription {
		fmt.Printf("Weapon Property Description: %v", weaponPropertDescription)
	}
	fmt.Printf("Cost: %v\n", weapon.cost)
	fmt.Printf("Weight: %v\n", weapon.weight)
	for damageType, damageAmount := range weapon.damage {
		fmt.Printf("Damage: %v Dice: %v\n", damageType, damageAmount)
	}
	fmt.Printf("Range Minimum: %v\n", weapon.rangemin)
	fmt.Printf("Range Maximum: %v\n", weapon.rangemax)
	fmt.Printf("Source: %v\n", weapon.source)
}

func (weapon *Weapon) GetItemName() string {
	return weapon.name
}

func (weapon *Weapon) GetItemTags() []string {
	return weapon.weaponProperty
}
