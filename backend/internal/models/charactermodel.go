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
	Skills           []map[string][]map[string]interface{}
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
	skillDependecies := map[string][]string{
		"Strength":     {"Atheletics"},
		"Dexterity":    {"Acrobatics", "Sleight of Hand", "Stealth"},
		"Constitution": {},
		"Intelligence": {"Arcana", "History", "Investigation", "Nature", "Religion"},
		"Wisdom":       {"Animal Handling", "Insight", "Medicine", "Perception", "Survival"},
		"Charisma":     {"Deception", "Intimidation", "Performance", "Persuasion"},
	}

	character.AbilityScores = make([]map[string]interface{}, len(abilityScoreNames))
	character.SavingThrows = make([]map[string]interface{}, len(abilityScoreNames))
	character.Skills = make([]map[string][]map[string]interface{}, len(abilityScoreNames))
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
		skills := skillDependecies[abilityScoreName]
		arrayOfAttributeSkills := generateArrayOfSkillMapPair(skills, abilityScoreStruct.abilityScoreModifier)
		attributeToSkillsMap := make(map[string][]map[string]interface{})
		attributeToSkillsMap[abilityScoreName] = arrayOfAttributeSkills
		character.Skills = append(character.Skills, attributeToSkillsMap)
	}

	return &character, nil
}
