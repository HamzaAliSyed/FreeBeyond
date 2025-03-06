package models

import (
	"backend/pkg/utils"
	"fmt"
)

type AbilityScore struct {
	abilityScoreValue    int
	abilityScoreModifier int
}

type SavingThrow struct {
	proficiencyBonus float64
	modifier         int
	otherbonus       int
	hasAdvantage     bool
	hasDisadvantage  bool
	total            int
}

type Skill struct {
	proficiencyBonus float64
	modifier         int
	otherbonus       int
	hasAdvantage     bool
	hasDisadvantage  bool
	total            int
}

func GenerateAbilityScoreStruct(isNew bool, score int) (AbilityScore, error) {
	var newAbilityScore AbilityScore
	var HigherConstant int
	if isNew {
		HigherConstant = 18
	} else {
		HigherConstant = 30
	}

	if score > HigherConstant || score <= 1 {
		return newAbilityScore, fmt.Errorf("cannot generate the ability score because of invalid score value")
	}

	modifierValue := utils.ModifierGenerator(score)
	newAbilityScore.abilityScoreValue = score
	newAbilityScore.abilityScoreModifier = modifierValue
	return newAbilityScore, nil
}

func (abilityScore AbilityScore) Print() {
	fmt.Println("Score: ", abilityScore.abilityScoreValue)
	fmt.Println("Modifier: ", abilityScore.abilityScoreModifier)
}

func GenerateSavingThrowForFirstTime(abilityScore AbilityScore) SavingThrow {
	var savingThrow SavingThrow
	savingThrow.hasAdvantage = false
	savingThrow.hasDisadvantage = false
	savingThrow.proficiencyBonus = 0
	savingThrow.otherbonus = 0
	savingThrow.modifier = abilityScore.abilityScoreModifier
	savingThrow.total = abilityScore.abilityScoreModifier

	return savingThrow
}

func (savingThrow SavingThrow) Print() {
	fmt.Println("Proficiency Bonus: ", savingThrow.proficiencyBonus)
	fmt.Println("Other Bonus: ", savingThrow.otherbonus)

	if savingThrow.hasAdvantage {
		fmt.Println("Has Advantage: Yes")
	} else {
		fmt.Println("Has Advantage: No")
	}

	if savingThrow.hasDisadvantage {
		fmt.Println("Has Disadvantage: Yes")
	} else {
		fmt.Println("Has Disadvantage: No")
	}

	fmt.Println("Modifier: ", savingThrow.modifier)
	fmt.Println("Total: ", savingThrow.total)
}

func generateNewSkillStructForTheFirstTime(modifier int) Skill {
	var newSkill Skill
	newSkill.hasAdvantage = false
	newSkill.hasDisadvantage = false
	newSkill.proficiencyBonus = 0
	newSkill.otherbonus = 0
	newSkill.modifier = modifier
	newSkill.total = modifier
	return newSkill
}

func (skill Skill) Print() {
	fmt.Println("Proficiency Bonus: ", skill.proficiencyBonus)
	fmt.Println("Other Bonus: ", skill.otherbonus)

	if skill.hasAdvantage {
		fmt.Println("Has Advantage: Yes")
	} else {
		fmt.Println("Has Advantage: No")
	}

	if skill.hasDisadvantage {
		fmt.Println("Has Disadvantage: Yes")
	} else {
		fmt.Println("Has Disadvantage: No")
	}

	fmt.Println("Modifier: ", skill.modifier)
	fmt.Println("Total: ", skill.total)
}

func generateSkillMapPair(skillName string, modifier int) map[string]interface{} {
	newSkill := make(map[string]interface{})
	newSkill[skillName] = generateNewSkillStructForTheFirstTime(modifier)
	return newSkill
}

func generateArrayOfSkillMapPair(skillNames []string, modifier int) []map[string]interface{} {
	skillsArray := make([]map[string]interface{}, 0)
	for _, skillName := range skillNames {
		skillMapPair := generateSkillMapPair(skillName, modifier)
		skillsArray = append(skillsArray, skillMapPair)
	}

	return skillsArray
}
