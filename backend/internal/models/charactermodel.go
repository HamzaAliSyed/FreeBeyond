package models

import (
	"backend/pkg/utils"
	"fmt"
)

type Character struct {
	Name             string
	ProficiencyBonus int
	AbilityScores    []map[string]interface{}
	SavingThrows     []map[string]interface{}
}

type CharacterParameters struct {
	Name          string
	AbilityScores []int
}

func NewCharacter(parameters CharacterParameters) (*Character, error) {
	validName, validNameError := utils.ValidateName(parameters.Name)

	var character Character

	if validNameError != nil {
		return nil, validNameError
	}

	character.Name = validName

	if len(parameters.AbilityScores) != 6 {
		return nil, fmt.Errorf("input array is not consistent")
	}

	character.ProficiencyBonus = 2

	abilityScoreNames := []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}
	character.AbilityScores = make([]map[string]interface{}, len(abilityScoreNames))
	character.SavingThrows = make([]map[string]interface{}, len(abilityScoreNames))
	for index, abilityScoreName := range abilityScoreNames {
		character.AbilityScores[index] = make(map[string]interface{})
		abilityScoreStruct, abilityScoreStructGenerationError := GenerateAbilityScoreStruct(true, parameters.AbilityScores[index])
		if abilityScoreStructGenerationError != nil {
			return nil, fmt.Errorf("error in generating character: %v", abilityScoreStructGenerationError)
		}
		character.AbilityScores[index][abilityScoreName] = abilityScoreStruct
		character.SavingThrows[index] = make(map[string]interface{})
		savingThrow := GenerateSavingThrowForFirstTime(abilityScoreStruct)
		character.SavingThrows[index][abilityScoreName] = savingThrow
	}

	return &character, nil
}
