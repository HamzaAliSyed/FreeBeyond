package models

import (
	"backend/pkg/utils"
	"fmt"
)

type AbilityScore struct {
	Name     string
	Score    int
	Modifier int
}

type AbilityScoreParameters struct {
	Name  string
	Score int
}

func NewAbilityScore(score AbilityScoreParameters) (*AbilityScore, error) {
	abilityScoreValue, abilityScoreModifer, abilityScoreError := utils.NewCharacterValidateScoreAndGenerateModifier(score.Score)
	if abilityScoreError != nil {
		return nil, fmt.Errorf("%v is not a valid value for %v", score.Score, score.Name)
	}

	return &AbilityScore{
		Name:     score.Name,
		Score:    abilityScoreValue,
		Modifier: abilityScoreModifer,
	}, nil
}
