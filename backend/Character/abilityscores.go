package character

import "fmt"

type AbilityScore struct {
	name     string `bson:"name"`
	score    int    `bson:"score"`
	modifier int    `bson:"modifier"`
}

func (abilityScore *AbilityScore) CreateAbilityScore(name string, value int) {
	abilityScore.name = name
	abilityScore.score = value
	abilityScore.modifier = (value - 10) / 2
}

func (abilityScore *AbilityScore) UpdateAbilityScore(value int) {
	abilityScore.score = value
	abilityScore.modifier = (value - 10) / 2
}

func (abilityScore *AbilityScore) String() string {
	return fmt.Sprintf("%s: Score %d, Modifier %d", abilityScore.name, abilityScore.score, abilityScore.modifier)
}
