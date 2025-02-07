package models

import (
	"backend/pkg/utils"
)

type Character struct {
	Name          string
	AbilityScores []AbilityScore
}

type CharacterParameters struct {
	Name                        string
	AbilityScoreParametersArray []AbilityScoreParameters
}

func NewCharacter(parameters CharacterParameters) (*Character, error) {
	validName, validNameError := utils.ValidateName(parameters.Name)
	if validNameError != nil {
		return nil, validNameError
	}

	var abilityScores []AbilityScore

	for _, abilityScoreParameter := range parameters.AbilityScoreParametersArray {
		abilityScore, abilityScoreCreationError := NewAbilityScore(abilityScoreParameter)
		if abilityScoreCreationError != nil {
			return nil, abilityScoreCreationError
		}

		abilityScores = append(abilityScores, *abilityScore)
	}

	return &Character{
		Name:          validName,
		AbilityScores: abilityScores,
	}, nil
}
