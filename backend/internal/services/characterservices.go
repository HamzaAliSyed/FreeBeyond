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

func (characterService CharacterService) CreateNewCharacter(characterparameters models.CharacterParameters) (models.Character, error) {
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

func (characterService CharacterService) PrintCharacterSheet(character models.Character) {
	fmt.Println("Printing Character Sheet")
	fmt.Printf("Character Name:%s\n", character.Name)
	fmt.Println("****************")
}
