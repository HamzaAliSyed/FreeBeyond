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
	fmt.Println("Printing Ability Scores")
	for abilityName, abilityScore := range character.AbilityScores {
		fmt.Println("****************")
		fmt.Println(abilityName)
		scoreStruct, ok := abilityScore.(models.AbilityScore)
		if ok {
			fmt.Println("Score: ", scoreStruct.Score)
			fmt.Println("Modifier: ", scoreStruct.Modifier)
		} else {
			fmt.Println("Cannot generate Ability Score Adequately")
		}
	}
	fmt.Println("****************")
	if len(character.Attacks) != 0 {
		for _, attack := range character.Attacks {
			attack.Print()
		}
	} else {
		fmt.Println("This character has no attacks")
	}
}

func (characterService CharacterService) UpdateAbilityScore(character *models.Character, abilityName string, newScore int) error {
	abilityScore, exists := character.AbilityScores[abilityName]
	if !exists {
		return fmt.Errorf("ability name %s does not exist", abilityName)
	}

	abilityScoreStruct, ok := abilityScore.(models.AbilityScore)
	if !ok {
		return fmt.Errorf("cannot parse the ability score struct")
	}

	score, modifier, updateError := utils.ValidateScoreAndGenerateModifier(newScore)
	if updateError != nil {
		return fmt.Errorf("error updating %v", updateError)
	}

	abilityScoreStruct.Score = score
	abilityScoreStruct.Modifier = modifier

	character.AbilityScores[abilityName] = abilityScoreStruct

	return nil
}

func (characterService CharacterService) IncreaseAbilityScore(character *models.Character, abilityName string, increaseAmount int) error {
	abilityScore, exists := character.AbilityScores[abilityName]
	if !exists {
		return fmt.Errorf("ability name %s does not exist", abilityName)
	}

	abilityScoreStruct, ok := abilityScore.(models.AbilityScore)
	if !ok {
		return fmt.Errorf("cannot parse the ability score struct")
	}

	newScore := abilityScoreStruct.Score + increaseAmount

	score, modifier, updateError := utils.ValidateScoreAndGenerateModifier(newScore)
	if updateError != nil {
		return fmt.Errorf("error updating %v", updateError)
	}

	abilityScoreStruct.Score = score
	abilityScoreStruct.Modifier = modifier

	character.AbilityScores[abilityName] = abilityScoreStruct

	return nil
}

func (characterService CharacterService) AddAttackToCharacter(character *models.Character, attackParameters models.AttackParameters) error {
	switch attackParameters.Type {
	case "ACBeatingAttack":
		{
			newAttack, newAttackGenerationError := models.GenerateACBeatingAttackForCharacter(character, attackParameters)
			if newAttackGenerationError != nil {
				return fmt.Errorf("encountered error in generating new Attack: %v", newAttackGenerationError)
			}
			character.Attacks = append(character.Attacks, newAttack)
		}
	default:
		return fmt.Errorf("invalid attack type")
	}
	return nil
}

func (characterService CharacterService) RemoveAttackFromCharacter(character *models.Character, name string) error {
	for index, attack := range character.Attacks {
		if attack.Name() == name {
			character.Attacks = append(character.Attacks[:index], character.Attacks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("attack with name %s not found", name)
}
