package main

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
)

func main() {
	characterService := services.NewCharacterService()

	characterParameters := models.CharacterParameters{Name: "Rewold Krushhammer"}
	character, characterCreateError := characterService.CreateNewCharacter(characterParameters)
	if characterCreateError != nil {
		fmt.Println("Error creating character:", characterCreateError)
		return
	}

	characterService.PrintCharacterSheet(character)

	characterNameUpdateError := characterService.UpdateCharacterName(character, "Harry 2")
	if characterNameUpdateError != nil {
		fmt.Println("Error updating character name:", characterNameUpdateError)
	}

	characterService.PrintCharacterSheet(character)

	characterNameUpdateErrorTwo := characterService.UpdateCharacterName(character, "Harry II")
	if characterNameUpdateErrorTwo != nil {
		fmt.Println("Error updating character name:", characterNameUpdateErrorTwo)
		return
	}

	characterService.PrintCharacterSheet(character)

	characterService.DeleteCharacterName(character)
	characterService.PrintCharacterSheet(character)

}
