package character

import "fmt"

type Skill struct {
	name                  string  `bson:"name"`
	abilityName           string  `bson:"skillused"`
	scoreModifier         int     `bson:"scoremodifier"`
	additionalBonus       int     `bson:"additionalbonus"`
	hasAdvantage          bool    `bson:"hasadvantage"`
	hasDisadvantage       bool    `bson:"hasDisadvantage"`
	numberOfProficiencies float64 `bson:"numberofproficiencies"`
	value                 int     `bson:"value"`
}

func (skill *Skill) CreateSkill(name, abilityName string, character Character) {
	skill.name = name
	skill.abilityName = abilityName
	modifier := character.CharacterAbilityScore(abilityName)
	skill.scoreModifier = modifier
	skill.additionalBonus = 0
	skill.numberOfProficiencies = 0
	skill.hasAdvantage = false
	skill.hasDisadvantage = false
	skill.value = modifier

}

func (skill *Skill) Print() {
	fmt.Printf("\nSkill Name: %v\nSkill Ability: %v\nAdvantage: %v \nDisadvantage: %v \nValue: %v\n", skill.name, skill.abilityName, skill.hasAdvantage, skill.hasDisadvantage, skill.value)
}
