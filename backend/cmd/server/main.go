package main

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
)

func main() {
	characterService := services.NewCharacterService()
	characterParameters := models.CharacterParameters{
		Name:          "Absol Curry",
		AbilityScores: []int{18, 18, 18, 18, 18, 18},
	}

	character, characterCreationError := characterService.CreateNewCharacter(characterParameters)
	if characterCreationError != nil {
		fmt.Printf("Cannot create the character %v", characterCreationError)
	}

	characterService.PrintCharacterSheet(character)
}
