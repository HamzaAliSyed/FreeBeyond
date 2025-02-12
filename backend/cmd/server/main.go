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
			{Name: "Strength", Score: 12},
			{Name: "Dexterity", Score: 11},
		},
	}
	character, characterCreateError := characterService.CreateNewCharacter(characterParameters)
	if characterCreateError != nil {
		fmt.Printf("error creating character: %v\n", characterCreateError)
		return
	}

	characterService.PrintCharacterSheet(character)
}
