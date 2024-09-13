package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
	id               primitive.ObjectID `bson:"_id"`
	characterName    string             `bson:"charactername"`
	proficiencyBonus int                `bson:"proficiencybonus"`
	abilityScores    []AbilityScore     `bson:"abilityscores"`
	savingThrows     []SavingThrow      `bson:"savingthrows"`
}

func (character *Character) GetID() primitive.ObjectID {
	return character.id
}

func (character *Character) GetIDasString() string {
	idString := character.GetID().Hex()
	return idString
}

func (character *Character) GetCharacterName() string {
	return character.characterName
}

func (character *Character) SetCharacterName(name string) {
	character.characterName = name
}

func (character *Character) SetCharacterAbilityScore(name string, value int) {
	var newAbilityScore AbilityScore
	newAbilityScore.CreateNewAbilityScore(name, value)
	character.abilityScores = append(character.abilityScores, newAbilityScore)
}

func (character *Character) GetCharacterAbilityScore(name string) (*AbilityScore, error) {
	for index, ability := range character.abilityScores {
		if ability.GetAbilityName() == name {
			return &character.abilityScores[index], nil
		}
	}
	return nil, errors.New("ability score not found")
}

func (character *Character) GetAllCharacterAbilityScores() []AbilityScore {
	return character.abilityScores
}

func (character *Character) UpdateAbilityScore(abilityName string, value int) error {
	abilityScore, abilityScoreRetrieveError := character.GetCharacterAbilityScore(abilityName)
	if abilityScoreRetrieveError != nil {
		return errors.New("cannot retrieve the ability score from the character")
	}

	newValue := abilityScore.GetAbilityScore() + value

	abilityScore.ImproveAbilityScore(newValue)

	for index, targetAbilityScore := range character.abilityScores {
		if targetAbilityScore.GetAbilityName() == abilityName {
			character.abilityScores[index] = *abilityScore
		}
	}

	savingThrow, savingThrowError := character.GetCharacterSavingThrow(abilityName)
	if savingThrowError != nil {
		return savingThrowError
	}

	savingThrow.SetSavingThrowMod(abilityScore.GetAbilityModifier())
	finalSavingThrowValue := savingThrow.CalculateValue(character)
	savingThrow.SetSavingThrowValue(finalSavingThrowValue)

	// Add skills later

	return nil
}

func (character *Character) SetProficiencyBonus(bonus int) {
	character.proficiencyBonus = bonus
}

func (character *Character) GetProficiencyBonus() int {
	return character.proficiencyBonus
}

func (character *Character) GetCharacterSavingThrows() []SavingThrow {
	return character.savingThrows
}

func (character *Character) CreateCharacterSavingThrow(name string) {
	abilityScore, _ := character.GetCharacterAbilityScore(name)
	mod := abilityScore.GetAbilityModifier()
	var savingThrow SavingThrow
	savingThrow.CreateSavingThrow(name, mod)
}

func (character *Character) GetCharacterSavingThrow(name string) (*SavingThrow, error) {
	for index, savingThrow := range character.GetCharacterSavingThrows() {
		if savingThrow.GetSavingThrowName() == name {
			return &character.savingThrows[index], nil
		}
	}

	return nil, errors.New("saving throw doesnt exist in character")
}
