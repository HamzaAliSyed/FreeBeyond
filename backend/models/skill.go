package models

type Skill struct {
	skillName       string  `bson:"skillname"`
	skillAtrribute  string  `bson:"skillattributevalue"`
	skillModValue   int     `bson:"skillmodvalue"`
	additionalBonus int     `bson:"additionalbonus"`
	hasAdvantage    bool    `bson:"hasadvantage"`
	hasDisadvantage bool    `bson:"hasdisadvantage"`
	proficiencies   float64 `bson:"proficiencies"`
	value           int     `bson:"value"`
}

func (skill *Skill) GetSkillName() string {
	return skill.skillName
}

func (skill *Skill) SetSkillName(skillName string) {
	skill.skillName = skillName
}

func (skill *Skill) GetSkillAttribute() string {
	return skill.skillAtrribute
}

func (skill *Skill) SetSkillAttribute(abilityScore string) {
	skill.skillAtrribute = abilityScore
}

func (skill *Skill) GetSkillModValue() int {
	return skill.skillModValue
}

func (skill *Skill) UpdateSkillModValue(character *Character) error {
	modSkill, modSkillError := character.GetCharacterAbilityScore(skill.skillAtrribute)
	if modSkillError != nil {
		return modSkillError
	}

	modValue := modSkill.GetAbilityModifier()
	skill.skillModValue = modValue
	return nil
}

func (skill *Skill) AddAdditionalBonus(value int) {
	skill.additionalBonus += value
}

func (skill *Skill) GetAdditionalBonus() int {
	return skill.additionalBonus
}

func (skill *Skill) ToggleSkillAdvantage(value bool) {
	skill.hasAdvantage = value
}
func (skill *Skill) ToggleSkillDisadvantage(value bool) {
	skill.hasDisadvantage = value
}

func (skill *Skill) DoesHasAdvantage() bool {
	return skill.hasAdvantage
}

func (skill *Skill) DoesHasDisadvantage() bool {
	return skill.hasDisadvantage
}

func (skill *Skill) AddProficiencies(value float64) {
	skill.proficiencies += value
}

func (skill *Skill) ShowProficiencies() float64 {
	return skill.proficiencies
}

func (skill *Skill) CalculateValue(character *Character) {
	skill.value = skill.skillModValue + skill.additionalBonus + int(((skill.proficiencies) * float64(character.GetProficiencyBonus())))
}

func (skill *Skill) GetValue() int {
	return skill.value
}

func (skill *Skill) CreateNewSkill(character *Character, skillname, skillatribute string) {
	skill.SetSkillName(skillname)
	skill.SetSkillAttribute(skillatribute)
	skill.UpdateSkillModValue(character)
	skill.additionalBonus = 0
	skill.hasAdvantage = false
	skill.hasDisadvantage = false
	skill.proficiencies = 0
	skill.CalculateValue(character)
}
