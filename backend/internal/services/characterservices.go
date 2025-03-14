package services

import (
	"backend/internal/models"
	"fmt"
)

type CharacterService struct{}

func (characterService CharacterService) GenerateNewCharacter(paramaters models.CharacterParameters) (*models.Character, error) {
	newCharacter, newCharacterError := models.CreateNewCharacter(paramaters)
	if newCharacterError != nil {
		return nil, fmt.Errorf("cannot create character: %v", newCharacterError)
	}

	return newCharacter, nil
}

func (characterService CharacterService) PrintCharacterSheet(character *models.Character) {
	fmt.Println("**********")
	fmt.Println("Character Name: ", character.GetCharacterName())
	fmt.Println("**********")
	fmt.Println("Proficiency Bonus: ", character.GetProficiencyBonus())
	character.PrintAbilityScores()
	fmt.Println("**********")
	character.PrintSavingThrows()
	fmt.Println("**********")
}

func (characterService CharacterService) ImproveAbilityScore(abilityName string, increase int, character models.Character) {
	character.IncreaseAbilityScore(abilityName, increase)
}
