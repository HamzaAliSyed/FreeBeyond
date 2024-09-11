package models

import (
	"backend/database"
	"backend/utils"

	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Damage string

const (
	Acid        Damage = "Acid"
	Bludgeoning Damage = "Bludgeoning"
	Cold        Damage = "Cold"
	Fire        Damage = "Fire"
	Force       Damage = "Force"
	Lightning   Damage = "Lightning"
	Necrotic    Damage = "Necrotic"
	Piercing    Damage = "Piercing"
	Poison      Damage = "Poison"
	Psychic     Damage = "Psychic"
	Radiant     Damage = "Radiant"
	Slashing    Damage = "Slashing"
	Thunder     Damage = "Thunder"
)

type Attacks interface {
	Create() interface{}
}

type ACBeatingAttacks struct {
	AttackName        string            `bson:"attackname"`
	DependentStat     string            `bson:"dependentstat"`
	DependentModifier int               `bson:"dependentmodifier"`
	IsProficienct     bool              `bson:"IsProficient"`
	RangeMin          int               `bson:"rangemin"`
	RangeMax          int               `bson:"rangemax"`
	TotalValue        int               `bson:"totalvalue"`
	Damage            map[Damage]string `bson:"damage"`
}

type DCSaveAttacks struct {
	AttackName string            `bson:"attackname"`
	SaveStat   string            `bson:"saveattack"`
	SaveDC     int               `bson:"savedc"`
	RangeMin   int               `bson:"rangemin"`
	RangeMax   int               `bson:"rangemax"`
	Damage     map[Damage]string `bson:"damage"`
}

func (acbeatingattacks *ACBeatingAttacks) Create() interface{} {
	return acbeatingattacks
}

func (dcsaveattacks *DCSaveAttacks) Create() interface{} {
	return dcsaveattacks
}

func NewACBeatingAttack(attackName, dependentStat string, rangeMin, rangeMax int, character *Character, damage map[Damage]string) *ACBeatingAttacks {
	var acbeatingattack ACBeatingAttacks
	acbeatingattack.AttackName = attackName
	acbeatingattack.DependentStat = dependentStat
	modifier := character.GetAbilityScoreModifier(dependentStat)
	acbeatingattack.DependentModifier = modifier
	acbeatingattack.RangeMin = rangeMin
	acbeatingattack.RangeMax = rangeMax
	acbeatingattack.Damage = damage
	acbeatingattack.TotalValue = modifier
	return &acbeatingattack
}
func NewDCSaveAttack(attackName, saveStat string, rangeMin, rangeMax int, character *Character, damage map[Damage]string) *DCSaveAttacks {
	var dcsaveattack DCSaveAttacks
	dcsaveattack.AttackName = attackName
	dcsaveattack.SaveStat = saveStat
	dcsaveattack.RangeMin = rangeMin
	dcsaveattack.RangeMax = rangeMax
	dcsaveattack.Damage = damage
	modifier := character.GetAbilityScoreModifier(saveStat)
	dcsaveattack.SaveDC = 8 + modifier
	return &dcsaveattack
}

type AbilityScore struct {
	StatName     string `bson:"statname"`
	StatValue    int    `bson:"statValue"`
	StatModifier int    `bson:"statmodifier"`
}

type SavingThrow struct {
	StatName              string  `bson:"statname"`
	SavingModifier        int     `bson:"savingmodifier"`
	AdditionalBonus       int     `bson:"additionalbonus"`
	NumberOfProficiencies float64 `bson:"numberofproficiencies"`
	HasAdvantage          bool    `bson:"hasadvantage"`
	HasDisadvantage       bool    `bson:"hasdisadvantage"`
	Value                 int     `bson:"value"`
}

type Skill struct {
	SkillName             string  `bson:"skillname"`
	StatAttribute         string  `bson:"statattribute"`
	StatModifier          int     `bson:"statmodifier"`
	AdditionalBonus       int     `bson:"additionalbonus"`
	NumberOfProficiencies float64 `bson:"numberofproficiencies"`
	HasAdvantage          bool    `bson:"hasadvantage"`
	HasDisadvantage       bool    `bson:"hasdisadvantage"`
	Value                 int     `bson:"value"`
}

type Character struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty"`
	charactername       string             `bson:"charactername,omitempty"`
	abilityscores       []AbilityScore     `bson:"abilityscores,omitempty"`
	savingthrows        []SavingThrow      `bson:"savingthrows,omitempty"`
	skills              []Skill            `bson:"skills,omitempty"`
	passives            []Skill            `bson:"passives,omitempty"`
	attacks             []Attacks          `bson:"attacks,omitempty"`
	weaponproficiencies []string           `bson:"weaponproficiencies"`
	inventory           map[*Item]int      `bson:"inventory"`
}

func (character Character) SetName(name string) Character {
	character.charactername = name
	return character
}

func (character *Character) GetName(id string) (string, error) {
	queryfilter := bson.M{
		"_id": id,
	}

	querySearchCharacterError := database.Characters.FindOne(context.TODO(), queryfilter).Decode(&character)
	if querySearchCharacterError != nil {
		return "", querySearchCharacterError
	}

	return character.charactername, nil

}

func CreateAbilityScore(statname string, statvalue int) (*AbilityScore, error) {
	var abilityscore AbilityScore
	if statvalue < 3 || statvalue > 18 {
		return nil, errors.New("invalid rolled stat, stat value must be between 3 and 18")
	}

	abilityscore.StatName = statname
	abilityscore.StatValue = statvalue
	abilityscore.StatModifier = (statvalue - 10) / 2
	return &abilityscore, nil
}

func CreateSavingThrow(statname string, statmodifier int) *SavingThrow {
	var savingthrow SavingThrow
	savingthrow.StatName = statname
	savingthrow.SavingModifier = statmodifier
	savingthrow.AdditionalBonus = 0
	savingthrow.NumberOfProficiencies = 0
	savingthrow.HasAdvantage = false
	savingthrow.HasDisadvantage = false
	savingthrow.Value = statmodifier + savingthrow.AdditionalBonus

	return &savingthrow
}

func CreateSkill(skillname, statattribute string, statmodifier int) *Skill {
	var skill Skill
	skill.SkillName = skillname
	skill.StatAttribute = statattribute
	skill.StatModifier = statmodifier
	skill.AdditionalBonus = 0
	skill.NumberOfProficiencies = 0
	skill.HasAdvantage = false
	skill.HasDisadvantage = false
	skill.Value = statmodifier + skill.AdditionalBonus
	return &skill
}

func CreatePassive(skillname, statattribute string, statmodifier int) *Skill {
	var passive Skill
	passive.SkillName = skillname
	passive.StatAttribute = statattribute
	passive.StatModifier = statmodifier
	passive.AdditionalBonus = 0
	passive.NumberOfProficiencies = 0
	passive.HasAdvantage = false
	passive.HasDisadvantage = false
	passive.Value = statmodifier + 10
	return &passive
}

func (character *Character) AddAbilityScoreToCharacter(abilityscore AbilityScore) {
	character.abilityscores = append(character.abilityscores, abilityscore)
}

func (character *Character) AddSavingThrowToCharacter(savingthrow SavingThrow) {
	character.savingthrows = append(character.savingthrows, savingthrow)
}

func (character *Character) AddSkillToCharacter(skill Skill) {
	character.skills = append(character.skills, skill)
}

func (character *Character) AddPassiveToCharacter(passive Skill) {
	character.passives = append(character.passives, passive)
}

func (character *Character) AddAttack(attack Attacks) {
	character.attacks = append(character.attacks, attack)
}

func (character *Character) AddWeaponProficiencies(weaponProficiency string) {
	if len(character.weaponproficiencies) == 0 {
		character.weaponproficiencies = append(character.weaponproficiencies, weaponProficiency)
	} else {
		exist := utils.Contains(character.weaponproficiencies, weaponProficiency)
		if !exist {
			character.weaponproficiencies = append(character.weaponproficiencies, weaponProficiency)
		}
	}
}

func (character *Character) AddItemsToInventory(item *Item) {
	if count, exist := character.inventory[item]; exist {
		character.inventory[item] = count + 1
	} else {
		character.inventory[item] = 1
	}

	itemProperties := (*item).GetAllProperties()
	name := itemProperties["name"].(string)
	rangemin := itemProperties["rangemin"].(int)
	rangemax := itemProperties["rangemax"].(int)
	damage := itemProperties["damage"].(map[Damage]string)

	if checkIfWeapon(*item) {
		switch checkWeaponType(*item) {
		case "Melee Weapon":
			{
				attack := NewACBeatingAttack(name, "Strength", rangemin, rangemax, character, damage)
				character.AddAttack(attack)
			}
		case "Ranged Weapon":
			{
				attack := NewACBeatingAttack(name, "Dexterity", rangemin, rangemax, character, damage)
				character.AddAttack(attack)
			}
		}
	}

}

func (character *Character) GetCharacterName() string {
	return character.charactername
}

func (character *Character) GetAllAbilityScore() []AbilityScore {
	return character.abilityscores
}

func (character *Character) GetAbilityScoreModifier(abilityscore string) int {
	abilityScores := character.GetAllAbilityScore()
	var Modifier int

	for index, abilityScore := range abilityScores {
		if abilityScore.StatName == abilityscore {
			Modifier = abilityScores[index].StatModifier
		}
	}

	return Modifier
}

func (character *Character) GetAllSavingThrow() []SavingThrow {
	return character.savingthrows
}

func (character *Character) GetAllSkills() []Skill {
	return character.skills
}

func (character *Character) GetAllPassives() []Skill {
	return character.passives
}

func (character *Character) GetAllAttacks() []Attacks {
	return character.attacks
}

func (character *Character) GetAllWeaponProficiencies() []string {
	return character.weaponproficiencies
}

func checkIfWeapon(item Item) bool {
	weaponClassifications := []string{"Simple Weapon", "Martial Weapon"}
	itemProperties := item.GetAllProperties()
	typetags := itemProperties["typetags"].([]string)
	for _, tag := range weaponClassifications {
		if utils.Contains(typetags, tag) {
			return true
		}
	}
	return false
}

func checkWeaponType(item Item) string {
	weaponType := []string{"Melee Weapon", "Ranged Weapon"}
	itemProperties := item.GetAllProperties()
	typetags := itemProperties["typetags"].([]string)
	for _, _type := range weaponType {
		if utils.Contains(typetags, _type) {
			return _type
		}
	}

	message := fmt.Errorf("not a weapon, something went wrong")
	return message.Error()
}
