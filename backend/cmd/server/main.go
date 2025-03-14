package main

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
)

func main() {
	characterService := services.CharacterService{}

	newCharacterParameters := models.CharacterParameters{
		Name:          "Absol Curry",
		AbilityScores: []int{18, 18, 18, 18, 18, 18},
	}

	character, characterCreationError := characterService.GenerateNewCharacter(newCharacterParameters)
	if characterCreationError != nil {
		fmt.Printf("Error: Cannot create new character: %v", characterCreationError)
		return
	}

	characterService.PrintCharacterSheet(character)

	fmt.Println("Absol goes into a library and study tirelessly for 72 hours")

	characterService.ImproveAbilityScore("Intelligence", 5, *character)
	characterService.PrintCharacterSheet(character)
}
