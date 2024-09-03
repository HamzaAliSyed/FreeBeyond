package utils

import (
	"backend/models"
	"fmt"
)

func IncreaseAbilityScore(character *models.Character, abilityScore string, value int) (*models.Character, error) {
	var updatedIndex int
	var found bool

	for index, score := range character.AbilityScores {
		if abilityScore == score.AbilityName {
			character.AbilityScores[index].Score += value
			character.AbilityScores[index].Modifier = (character.AbilityScores[index].Score - 10) / 2
			updatedIndex = index
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ability score not found")
	}

	for index, save := range character.SavingThrows {
		if save.AttributeName == abilityScore {
			character.SavingThrows[index].Modifer = character.AbilityScores[updatedIndex].Modifier
			character.SavingThrows[index].Value = save.AdditionalBonus + character.SavingThrows[index].Modifer + (int(save.NumberOProficiency) * character.ProficiencyBonus)
		}
	}

	for i, skill := range character.Skills {
		if GetAbilityForSkill(skill.AttributeName) == abilityScore {
			character.Skills[i].Modifer = character.AbilityScores[updatedIndex].Modifier
			character.Skills[i].Value = character.Skills[i].Modifer + character.Skills[i].AdditionalBonus
		}
	}

	return character, nil
}

func GetAbilityForSkill(skillName string) string {
	skillToAbility := map[string]string{
		"Acrobatics":      "Dexterity",
		"Animal Handling": "Wisdom",
		"Arcana":          "Intelligence",
		"Athletics":       "Strength",
		"Deception":       "Charisma",
		"History":         "Intelligence",
		"Insight":         "Wisdom",
		"Intimidation":    "Charisma",
		"Investigation":   "Intelligence",
		"Medicine":        "Wisdom",
		"Nature":          "Intelligence",
		"Perception":      "Wisdom",
		"Performance":     "Charisma",
		"Persuasion":      "Charisma",
		"Religion":        "Intelligence",
		"Sleight of Hand": "Dexterity",
		"Stealth":         "Dexterity",
		"Survival":        "Wisdom",
	}

	return skillToAbility[skillName]
}
