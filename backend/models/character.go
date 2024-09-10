package models

import (
	"backend/database"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterOperations interface{}

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
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	charactername string             `bson:"charactername,omitempty"`
	abilityscores []AbilityScore     `bson:"abilityscores,omitempty"`
	savingthrows  []SavingThrow      `bson:"savingthrows,omitempty"`
	skills        []Skill            `bson:"skills"`
	passives      []Skill            `bson:"passives"`
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
