package services

import (
	"backend/internal/models"
	"backend/pkg/utils"
	"fmt"
)

type CharacterService struct{}

func NewCharacterService() *CharacterService {
	return &CharacterService{}
}

func (characterService CharacterService) CreateNewCharacter(characterparameters models.CharacterParameters) (*models.Character, error) {
	return models.NewCharacter(characterparameters)
}

func (characterService CharacterService) UpdateCharacterName(character *models.Character, newName string) error {
	validName, validNameError := utils.ValidateName(newName)
	if validNameError != nil {
		return validNameError
	}

	character.Name = validName
	return nil
}

func (characterService CharacterService) DeleteCharacterName(character *models.Character) {
	character.Name = "Unnamed Micmansion"
}

func (characterService CharacterService) UpdateAbilityScore(character *models.Character, abilityName string, newScore int) {
	for index, ability := range character.AbilityScores {
		if ability.Name == abilityName {
			newScore, newModifierValue, updateError := utils.ValidateScoreAndGenerateModifier(newScore)
			if updateError != nil {
				fmt.Printf("encounter error in updating character: %v\n", updateError)
			}

			ability.Score = newScore
			ability.Modifier = newModifierValue

			character.AbilityScores[index] = ability

			fmt.Printf("THe new character ability %s have the score %d and modifier %d", ability.Name, ability.Score, ability.Modifier)
		}
	}
}

func (characterService CharacterService) PrintCharacterSheet(character *models.Character) {
	fmt.Println("Printing Character Sheet")
	fmt.Printf("Character Name:%s\n", character.Name)
	fmt.Println("Printing Ability Scores")
	for _, abilityScore := range character.AbilityScores {
		fmt.Printf("Ability Score: %s\n", abilityScore.Name)
		fmt.Printf("Value: %d\n", abilityScore.Score)
		fmt.Printf("Modifier: %d\n", abilityScore.Modifier)
		fmt.Println()
	}
}
