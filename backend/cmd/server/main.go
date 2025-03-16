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
	fmt.Println("And he goes to cardio and stretching a little bit")

	characterService.ImproveAbilityScore("Intelligence", 5, *character)
	characterService.ImproveAbilityScore("Dexterity", 10, *character)
	characterService.PrintCharacterSheet(character)
}
