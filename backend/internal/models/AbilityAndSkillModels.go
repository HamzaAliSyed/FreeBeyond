package models

import "fmt"

type AbilityScore struct {
	value    int
	modifier int
}

func (abilityScore AbilityScore) GetValue() int {
	return abilityScore.value
}

func (abilityScore AbilityScore) GetModifier() int {
	return abilityScore.modifier
}

func (abilityScore AbilityScore) Print() {
	fmt.Println("Value: ", abilityScore.value)
	fmt.Println("Modifier: ", abilityScore.modifier)
}

type SavingThrow struct {
	numberOfProficiencies float64
	hasAdvantage          bool
	hasDisadvantage       bool
	otherBonus            int
	proficiencyBonus      *int
	modifier              *int
}

func GenerateSavingThrow(savingThrowName string, character Character) SavingThrow {
	var savingThrow SavingThrow

	savingThrow.numberOfProficiencies = 0
	savingThrow.hasAdvantage = false
	savingThrow.hasDisadvantage = false
	savingThrow.otherBonus = 0
	savingThrow.proficiencyBonus = &character.proficiencyBonus
	savingThrow.modifier = &character.abilityScores[savingThrowName].modifier

	return savingThrow
}

func (savingThrow SavingThrow) TotalRoll() int {
	var total int
	scoreFromProficieny := int(savingThrow.numberOfProficiencies * float64(*savingThrow.proficiencyBonus))
	total = scoreFromProficieny + *savingThrow.modifier + savingThrow.otherBonus
	return total
}

func (savingThrow SavingThrow) Print() {
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

	fmt.Println("Total: ", savingThrow.TotalRoll())

}

type Skill struct {
	numberOfProficiencies float64
	hasAdvantage          bool
	hasDisadvantage       bool
	otherBonus            int
	proficiencyBonus      *int
	modifier              *int
}

func GenerateSkill(abilityScoreName string, character Character) Skill {
	var skill Skill
	skill.hasAdvantage = false
	skill.hasDisadvantage = false
	skill.otherBonus = 0
	skill.numberOfProficiencies = 0
	skill.proficiencyBonus = &character.proficiencyBonus
	skill.modifier = &character.abilityScores[abilityScoreName].modifier

	return skill
}

func (skill Skill) TotalRoll() int {
	var total int
	scoreFromProficieny := int(skill.numberOfProficiencies * float64(*skill.proficiencyBonus))
	total = scoreFromProficieny + *skill.modifier + skill.otherBonus
	return total
}
func (skill Skill) Print() {
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

	fmt.Println("Total: ", skill.TotalRoll())
}
