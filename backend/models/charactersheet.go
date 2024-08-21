package models

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	Name                 string             `bson:"name"`
	Strength             int                `bson:"strength"`
	Dexterity            int                `bson:"dexterity"`
	Constitution         int                `bson:"constitution"`
	Intelligence         int                `bson:"intelligence"`
	Wisdom               int                `bson:"wisdom"`
	Charisma             int                `bson:"charisma"`
	StrengthModifier     int                `bson:"strength_modifier"`
	DexterityModifier    int                `bson:"dexterity_modifier"`
	ConstitutionModifier int                `bson:"constitution_modifier"`
	IntelligenceModifier int                `bson:"intelligence_modifier"`
	WisdomModifier       int                `bson:"wisdom_modifier"`
	CharismaModifier     int                `bson:"charisma_modifier"`
	MaxCarryWeight       int                `bson:"maxcarryweight"`
	HitDie               string             `bson:"hitdie"`
	NumberOfHitdie       map[string]int     `bson:"numberofhitdie"`
	HitPoints            int                `bson:"hitpoints"`
	TempHitPoints        int                `bson:"temphitpoints"`
	Initiative           Initiative         `bson:"initiative"`
	ArmorClass           int                `bson:"armorclass"`
	CanDoSpellCasting    bool               `bson:"candospellcasting"`
	TypeOfSpellCasting   string             `bson:"typesofspellcasting"`
	SavingThrows         []SavingThrow      `bson:"savingthrows"`
	ProficiencyBonus     int                `bson:"proficiencybonus"`
	Skills               []Skill            `bson:"skills"`
	Class                map[string]int     `bson:"class"`
	WeaponProficiencies  []string           `bson:"weaponproficiencies"`
	ArmorProficiencies   []string           `bson:"armorproficiencies"`
}

type SavingThrow struct {
	Name                  string
	Attribute             string
	AttributeModifier     int
	NumberOfProficiencies float32
	HasAdvantage          bool
	HasDisvantage         bool
	Value                 int
}

type Skill struct {
	Name                  string
	Attribute             string
	AttributeModifier     int
	NumberOfProficiencies float32
	HasAdvantage          bool
	HasDisvantage         bool
	AdditionalBoost       int
	Value                 int
}

type Initiative struct {
	AdditionalBonus int
	FinalValue      int
}

func (character *Character) SetName(name string) {
	character.Name = name
}

func (character *Character) SetStrengthAbilityScore(score int) {
	character.Strength = score
	character.StrengthModifier = CalculateModifier(score)
	character.MaxCarryWeight = score * 15
}

func (character *Character) SetDexterityAbilityScore(score int) {
	character.Dexterity = score
	character.DexterityModifier = CalculateModifier(score)
}

func (character *Character) SetConstitutionAbilityScore(score int) {
	character.Constitution = score
	character.ConstitutionModifier = CalculateModifier(score)
}

func (character *Character) SetIntelligenceAbilityScore(score int) {
	character.Intelligence = score
	character.IntelligenceModifier = CalculateModifier(score)
}

func (character *Character) SetWisdomAbilityScore(score int) {
	character.Wisdom = score
	character.WisdomModifier = CalculateModifier(score)
}

func (character *Character) SetCharismaAbilityScore(score int) {
	character.Charisma = score
	character.CharismaModifier = CalculateModifier(score)
}

func (character *Character) InitializeSavingThrows() {
	Attributes := []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}
	for _, attribute := range Attributes {
		var savingthrow SavingThrow
		savingthrow.Name = attribute
		savingthrow.Attribute = attribute
		switch attribute {
		case "Strength":
			{
				savingthrow.AttributeModifier = character.StrengthModifier
			}
		case "Dexterity":
			{
				savingthrow.AttributeModifier = character.DexterityModifier
			}
		case "Constitution":
			{
				savingthrow.AttributeModifier = character.ConstitutionModifier
			}
		case "Intelligence":
			{
				savingthrow.AttributeModifier = character.IntelligenceModifier
			}
		case "Wisdom":
			{
				savingthrow.AttributeModifier = character.WisdomModifier
			}
		case "Charisma":
			{
				savingthrow.AttributeModifier = character.CharismaModifier
			}
		}
		character.SavingThrows = append(character.SavingThrows, savingthrow)
	}
}

func (character *Character) AddProficiencyToSavingThrow(savingThrowName string, proficiencyValue float32) {
	for i, st := range character.SavingThrows {
		if st.Name == savingThrowName {
			character.SavingThrows[i].NumberOfProficiencies += proficiencyValue
			character.SavingThrows[i].Value = st.AttributeModifier + (int(character.SavingThrows[i].NumberOfProficiencies) * character.ProficiencyBonus)
			break
		}
	}
}

func (character *Character) InitializeCharacterSkill() {
	skillMap := make(map[string]string)
	skillMap["Acrobatics"] = "Dexterity"
	skillMap["Animal Handling"] = "Wisdom"
	skillMap["Arcana"] = "Intelligence"
	skillMap["Athletics"] = "Strength"
	skillMap["Deception"] = "Charisma"
	skillMap["History"] = "Intelligence"
	skillMap["Insight"] = "Wisdom"
	skillMap["Intimidation"] = "Charisma"
	skillMap["Investigation"] = "Intelligence"
	skillMap["Medicine"] = "Wisdom"
	skillMap["Nature"] = "Intelligence"
	skillMap["Perception"] = "Wisdom"
	skillMap["Performance"] = "Charisma"
	skillMap["Persuasion"] = "Charisma"
	skillMap["Religion"] = "Intelligence"
	skillMap["Sleight of Hand"] = "Dexterity"
	skillMap["Stealth"] = "Dexterity"
	skillMap["Survival"] = "Wisdom"

	for key, value := range skillMap {
		var skill Skill
		skill.Name = key
		skill.Attribute = value
		switch skill.Attribute {
		case "Strength":
			skill.AttributeModifier = character.StrengthModifier
		case "Dexterity":
			skill.AttributeModifier = character.DexterityModifier
		case "Constitution":
			skill.AttributeModifier = character.ConstitutionModifier
		case "Intelligence":
			skill.AttributeModifier = character.IntelligenceModifier
		case "Wisdom":
			skill.AttributeModifier = character.WisdomModifier
		case "Charisma":
			skill.AttributeModifier = character.CharismaModifier
		}
		character.Skills = append(character.Skills, skill)
	}
}

func (character *Character) AddProficiencyToSkill(skillName string, proficiencyValue float32) {
	for i, skill := range character.Skills {
		if skill.Name == skillName {
			skill.NumberOfProficiencies = proficiencyValue
			skill.Value = skill.AttributeModifier + skill.AdditionalBoost + (int(skill.NumberOfProficiencies) * character.ProficiencyBonus)
			character.Skills[i] = skill
			break
		}
	}
}

func (character *Character) FindHighestClassLevel() int {
	maxValue := 0

	for _, value := range character.Class {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func (character *Character) AbilityScoreImprovement(ability string, score int) error {

	corestats := []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}

	validAbility := false
	for _, stat := range corestats {
		if ability == stat {
			validAbility = true
			break
		}
	}
	if !validAbility {
		return fmt.Errorf("invalid ability: %s", ability)
	}

	switch ability {
	case "Strength":
		character.Strength = score
		character.StrengthModifier = CalculateModifier(score)
	case "Dexterity":
		character.Dexterity = score
		character.DexterityModifier = CalculateModifier(score)
	case "Constitution":
		character.Constitution = score
		character.ConstitutionModifier = CalculateModifier(score)
	case "Intelligence":
		character.Intelligence = score
		character.IntelligenceModifier = CalculateModifier(score)
	case "Wisdom":
		character.Wisdom = score
		character.WisdomModifier = CalculateModifier(score)
	case "Charisma":
		character.Charisma = score
		character.CharismaModifier = CalculateModifier(score)
	default:
		// Handle invalid ability names
		fmt.Printf("Invalid ability: %s\n", ability)
	}

	for i, savingthrow := range character.SavingThrows {
		if savingthrow.Attribute == ability {
			switch ability {
			case "Strength":
				character.SavingThrows[i].AttributeModifier = character.StrengthModifier
			case "Dexterity":
				character.SavingThrows[i].AttributeModifier = character.DexterityModifier
			case "Constitution":
				character.SavingThrows[i].AttributeModifier = character.ConstitutionModifier
			case "Intelligence":
				character.SavingThrows[i].AttributeModifier = character.IntelligenceModifier
			case "Wisdom":
				character.SavingThrows[i].AttributeModifier = character.WisdomModifier
			case "Charisma":
				character.SavingThrows[i].AttributeModifier = character.CharismaModifier
			}

			character.SavingThrows[i].Value = character.SavingThrows[i].AttributeModifier + (int(character.SavingThrows[i].NumberOfProficiencies) * character.ProficiencyBonus)
		}
	}

	for i, skill := range character.Skills {
		if skill.Attribute == ability {
			switch ability {
			case "Strength":
				character.Skills[i].AttributeModifier = character.StrengthModifier
			case "Dexterity":
				character.Skills[i].AttributeModifier = character.DexterityModifier
			case "Constitution":
				character.Skills[i].AttributeModifier = character.ConstitutionModifier
			case "Intelligence":
				character.Skills[i].AttributeModifier = character.IntelligenceModifier
			case "Wisdom":
				character.Skills[i].AttributeModifier = character.WisdomModifier
			case "Charisma":
				character.Skills[i].AttributeModifier = character.CharismaModifier
			}
			character.Skills[i].Value = character.Skills[i].AttributeModifier + character.Skills[i].AdditionalBoost + (int(character.Skills[i].NumberOfProficiencies) * character.ProficiencyBonus)
		}
	}

	return nil

}

func (character *Character) AddWeaponProficiency(weapontype ...string) {
	character.WeaponProficiencies = append(character.WeaponProficiencies, weapontype...)
}

func (character *Character) AddArmorProficiency(armortype ...string) {
	character.WeaponProficiencies = append(character.WeaponProficiencies, armortype...)
}

func (character *Character) SetInitiative(additionalboost int) {
	character.Initiative.AdditionalBonus = additionalboost
	character.Initiative.FinalValue = character.DexterityModifier + character.Initiative.AdditionalBonus
}

func (character *Character) SetArmorClass() {
	character.ArmorClass = 10 + character.DexterityModifier
}

func CalculateModifier(score int) int {
	mod := (score - 10) / 2
	return mod
}

func CompareHitDie(existinghitdie, newhitdie string) string {
	existingvalue := existinghitdie[2:]
	newvalue := newhitdie[2:]

	intexistingvalue, err1 := strconv.Atoi(existingvalue)
	intnewvalue, err2 := strconv.Atoi(newvalue)

	if err1 != nil || err2 != nil {
		return "Error: Invalid hit die value"
	}

	var largerValue int

	if intexistingvalue > intnewvalue {
		largerValue = intexistingvalue
	}

	if intnewvalue > intexistingvalue {
		largerValue = intnewvalue
	}

	hitdie := fmt.Sprintf("1d%v", largerValue)
	return hitdie

}

func (character *Character) SetProficiencyBonus() {
	ProficiencyTable := map[int]int{
		1:  2,
		2:  2,
		3:  2,
		4:  2,
		5:  3,
		6:  3,
		7:  3,
		8:  3,
		9:  4,
		10: 4,
		11: 4,
		12: 4,
		13: 5,
		14: 5,
		15: 5,
		16: 5,
		17: 6,
		18: 6,
		19: 6,
		20: 6,
	}

	highestlevel := character.FindHighestClassLevel()

	if bonus, exists := ProficiencyTable[highestlevel]; exists {
		character.ProficiencyBonus = bonus
	} else {
		character.ProficiencyBonus = 2
		fmt.Printf("Warning: Level %d not found in proficiency table.\n", highestlevel)
	}
}

func (character *Character) GenericAttributeModifier(attribute string) {
	switch attribute {

	}
}
