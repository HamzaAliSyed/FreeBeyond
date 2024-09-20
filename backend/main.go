package main

import (
	character "backend/Character"
	item "backend/Item"
	"fmt"
)

func main() {
	var HazaristSchaffenhifer character.Character

	HazaristSchaffenhifer.SetCharacterName("Hazarist Schaffenhifer")
	HazaristSchaffenhifer.SetProficiencyBonus(2)

	AbilityScores := map[string]int{
		"Strength":     15,
		"Dexterity":    16,
		"Constitution": 20,
		"Intelligence": 24,
		"Wisdom":       18,
		"Charisma":     14,
	}

	for abilityName, abilityScore := range AbilityScores {
		HazaristSchaffenhifer.AddAbilityScore(abilityName, abilityScore)
	}

	HazaristSchaffenhifer.PrintCharacterSheet()

	params := item.WeaponParams{
		Name:           "Longsword",
		Rarity:         item.Common,
		Description:    "A longsword is a versatile, double-edged weapon characterized by its straight, extended blade, typically around 35 to 45 inches in length. \nDesigned for both cutting and thrusting, it is wielded with two hands, offering a balance between speed and power. \nIt is used by knights and adventures all over the Faerun and excels in close quarters combat.",
		TypeTags:       []string{"Martial Weapon", "Melee Weapon"},
		WeaponProperty: []string{"Versatile"},
		DamageType:     []item.Damage{item.Slashing},
		DamageAmount:   []string{"1d8"},
		Cost:           "15 gp",
		Weight:         "3 lbs",
		Source:         "Player's Handbook",
		RangeMin:       0,
		RangeMax:       5,
	}

	Longsword, LongswordCreationError := item.CreateItemFactory("weapon", params)

	if LongswordCreationError != nil {
		fmt.Println(LongswordCreationError.Error())
		return
	}

	Longsword.Print()

}
