package models

import (
	"backend/pkg/utils"
	"fmt"
)

type Character struct {
	name             string
	ac               AC
	proficiencyBonus int
	abilityScores    map[string]*AbilityScore
	savingThrows     map[string]*SavingThrow
	skills           map[string]*Skill
}

type CharacterParameters struct {
	Name          string
	AbilityScores []int
}

func CreateNewCharacter(parameters CharacterParameters) (*Character, error) {
	validName, validNameError := utils.ValidateName(parameters.Name)
	if validNameError != nil {
		return nil, fmt.Errorf("name could not be validated %v", validNameError)
	}

	if len(parameters.AbilityScores) == 0 {
		return nil, fmt.Errorf("paramters are absent in the creation")
	}

	if len(parameters.AbilityScores) > 6 {
		return nil, fmt.Errorf("more than 6 parameters present")
	}

	var character Character
	character.abilityScores = make(map[string]*AbilityScore)
	character.savingThrows = make(map[string]*SavingThrow)
	character.skills = make(map[string]*Skill)
	character.name = validName
	character.proficiencyBonus = 2

	var mainAbilityNames = []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}

	for index, mainAbilityName := range mainAbilityNames {
		abilityStruct := &AbilityScore{
			value:    parameters.AbilityScores[index],
			modifier: utils.ModifierGenerator(parameters.AbilityScores[index]),
		}
		character.abilityScores[mainAbilityName] = abilityStruct
		savingThrowStruct := GenerateSavingThrow(mainAbilityName, character)
		character.savingThrows[mainAbilityName] = &savingThrowStruct
	}

	mainSkillToAttribute := map[string][]string{
		"Strength":     {"Athletics"},
		"Dexterity":    {"Acrobatics", "Sleight of Hand", "Stealth"},
		"Constitution": {},
		"Intelligence": {"Arcana", "History", "Investigation", "Nature", "Religion"},
		"Wisdom":       {"Animal Handling", "Insight", "Medicine", "Perception", "Survival"},
		"Charisma":     {"Deception", "Intimidation", "Performance", "Persuasion"},
	}

	for ability, skills := range mainSkillToAttribute {
		for _, skillName := range skills {
			skillStruct := GenerateSkill(ability, character)
			character.skills[skillName] = &skillStruct
		}
	}

	var ac AC
	ac.CalculateAC(character.abilityScores["Dexterity"].modifier)
	character.ac = ac

	return &character, nil

}

func (character Character) GetCharacterName() string {
	return character.name
}

func (character Character) PrintAbilityScores() {
	fmt.Println("Ability Scores")
	for name, abilityScore := range character.abilityScores {
		fmt.Println("Ability: ", name)
		abilityScore.Print()
	}
}

func (character Character) UpdateProficiencyBonus(newValue int) {
	character.proficiencyBonus = newValue
	fmt.Println("The new proficiency bonus is updated to: ", character.proficiencyBonus)
}

func (character Character) GetProficiencyBonus() int {
	return character.proficiencyBonus
}

func (character Character) PrintSavingThrows() {
	fmt.Println("Saving Throws")
	for name, savingThrow := range character.savingThrows {
		fmt.Println("Saving Throw: ", name)
		savingThrow.Print()
	}
}

func (character Character) PrintSkills() {
	fmt.Println("Skills")
	for name, skill := range character.skills {
		fmt.Println("Skill Name: ", name)
		skill.Print()
	}
}

func (character Character) IncreaseAbilityScore(ability string, value int) {
	oldValue := character.abilityScores[ability].value
	newValue := oldValue + value
	newMod := utils.ModifierGenerator(newValue)
	character.abilityScores[ability].value = newValue
	character.abilityScores[ability].modifier = newMod
}

func (character Character) GetAC() int {
	return character.ac.finalValue
}
