package character

import (
	"errors"
	"fmt"
)

type AbilityScore struct {
	name                 string `bson:"name"`
	abilityScoreValue    int    `bson:"abilityscorevalue"`
	abilityScoreModifier int    `bson:"abilityscoremodifier"`
}

func (abilityScore *AbilityScore) CreateAbilityScore(AS string, ASV int) error {
	if ASV < 3 || ASV > 18 {
		return errors.New("ability score value cannot be lower than 3 or greater than 18 when stats are rolled")
	}
	abilityScore.name = AS
	abilityScore.abilityScoreValue = ASV
	abilityScore.abilityScoreModifier = (ASV - 10) / 2
	return nil
}

func (abilityScore *AbilityScore) Print() {
	fmt.Printf("\nAbility Score: %v \n Value: %v \n Modifier: %v \n", abilityScore.name, abilityScore.abilityScoreValue, abilityScore.abilityScoreModifier)
}

func (abilityScore *AbilityScore) GetNameandMod() (string, int) {
	return abilityScore.name, abilityScore.abilityScoreModifier
}
