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
	character, characterGenerationError := models.NewCharacter(characterparameters)
	if characterGenerationError != nil {
		return models.Character{}, characterGenerationError
	}

	return *character, nil
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
	if character.Name == "" && len(character.Name) == 0 {
		fmt.Println("\nCannot Print cause character is empty")
		return
	}
	fmt.Println("\nPrinting Character Sheet")
	fmt.Println("****************")
	fmt.Println("****************")
	fmt.Printf("Character Name:%s\n", character.Name)
	fmt.Println("****************")
	fmt.Println("Ability Scores")
	fmt.Println("****************")
	fmt.Println("****************")
	for _, abilityScore := range character.AbilityScores {
		for abilityName, abilityStruct := range abilityScore {
			fmt.Println("Name: ", abilityName)
			structAS := abilityStruct.(models.AbilityScore)
			structAS.Print()
		}
	}
	fmt.Println("****************")
	fmt.Println("Saving Throws")
	fmt.Println("****************")
	for _, savingThrow := range character.SavingThrows {
		for savingThrowName, savingThrowStruct := range savingThrow {
			fmt.Println("Name: ", savingThrowName)
			savingThrowDecoded := savingThrowStruct.(models.SavingThrow)
			savingThrowDecoded.Print()
		}
		fmt.Println("****************")
	}
	fmt.Println("****************")
}
