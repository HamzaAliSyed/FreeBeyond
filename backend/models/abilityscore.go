package models

import "errors"

type AbilityScore struct {
	abilityname     string `bson:"abilityname"`
	abilityscore    int    `bson:"abilityscore"`
	abilitymodifier int    `bson:"abilitymodifier"`
}

func (abilityScore *AbilityScore) CreateNewAbilityScore(name string, score int) error {
	abilityScore.abilityname = name

	if score < 3 || score > 18 {
		return errors.New("ability score cannot be rolled to be higher than 18 or lower than 3")
	}

	abilityScore.abilityscore = score
	abilityScore.abilitymodifier = (score - 10) / 2

	return nil
}

func (abilityScore *AbilityScore) ImproveAbilityScore(value int) error {
	if value > 22 {
		return errors.New("ability score cannot increase 22")
	}

	abilityScore.abilityscore = value
	abilityScore.abilitymodifier = (value - 10) / 2
	return nil
}

func (abilityScore *AbilityScore) GetAbilityScore() int {
	return abilityScore.abilityscore
}

func (abilityScore *AbilityScore) GetAbilityModifier() int {
	return abilityScore.abilitymodifier
}

func (abilityScore *AbilityScore) GetAbilityName() string {
	return abilityScore.abilityname
}
