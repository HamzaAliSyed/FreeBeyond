package main

import (
	character "backend/Character"
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
}
