package main

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
)

func main() {
	characterService := services.NewCharacterService()

	characterParameters := models.CharacterParameters{
		Name: "Rewold Krushhammer",
		AbilityScoresParametersArray: []models.AbilityScoreParameters{
			{Name: "Strength", Score: 16},
			{Name: "Dexterity", Score: 14},
			{Name: "Constitution", Score: 15},
			{Name: "Intelligence", Score: 12},
			{Name: "Wisdom", Score: 13},
			{Name: "Charisma", Score: 10},
		},
	}

	character, characterCreateError := characterService.CreateNewCharacter(characterParameters)
	if characterCreateError != nil {
		fmt.Printf("error creating character: %v\n", characterCreateError)
		return
	}

	characterService.PrintCharacterSheet(character)

	characterUpdateError := characterService.UpdateAbilityScore(&character, "Dexterity", 16)
	if characterUpdateError != nil {
		fmt.Printf("Error Updating Stats: %v\n", characterUpdateError)
		return
	}

	characterService.PrintCharacterSheet(character)

	characterIncreaseError := characterService.IncreaseAbilityScore(&character, "Dexterity", 2)
	if characterIncreaseError != nil {
		fmt.Printf("Error in increasing character stats: %v\n", characterIncreaseError)
		return
	}

	characterService.PrintCharacterSheet(character)

}
