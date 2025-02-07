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
		AbilityScoreParametersArray: []models.AbilityScoreParameters{
			{Name: "Strength", Score: 18},
			{Name: "Dexterity", Score: 13},
			{Name: "Constitution", Score: 15},
			{Name: "Intelligence", Score: 9},
			{Name: "WIsdom", Score: 14},
			{Name: "Charisma", Score: 12},
		},
	}
	character, characterCreateError := characterService.CreateNewCharacter(characterParameters)
	if characterCreateError != nil {
		fmt.Println("Error creating character:", characterCreateError)
		return
	}

	characterService.PrintCharacterSheet(character)
}
