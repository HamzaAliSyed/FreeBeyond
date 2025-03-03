package models

import (
	"backend/pkg/utils"
	"fmt"
)

type Character struct {
	Name          string
	AbilityScores map[string]interface{}
}

type CharacterParameters struct {
	Name                         string
	AbilityScoresParametersArray []AbilityScoreParameters
}

func NewCharacter(parameters CharacterParameters) (Character, error) {
	validName, validNameError := utils.ValidateName(parameters.Name)
	var character Character
	character.AbilityScores = make(map[string]interface{})
	if validNameError != nil {
		return character, validNameError
	}

	for _, abilityScore := range parameters.AbilityScoresParametersArray {
		var characterAbility AbilityScore
		asi, asm, abilityScoreUpdateError := utils.NewCharacterValidateScoreAndGenerateModifier(abilityScore.Score)
		if abilityScoreUpdateError != nil {
			fmt.Printf("cannot create attribute score %v", abilityScoreUpdateError)
			return character, abilityScoreUpdateError
		}

		characterAbility.Name = abilityScore.Name
		characterAbility.Score = asi
		characterAbility.Modifier = asm
		character.AbilityScores[abilityScore.Name] = characterAbility
	}
	character.Name = validName
	return character, nil
}
