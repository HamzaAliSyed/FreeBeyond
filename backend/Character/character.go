package character

import (
	"fmt"
	"reflect"
)

type Character struct {
	name             string         `bson:"name"`
	proficiencyBonus int            `bson:"proficiencybonus"`
	abilityScores    []AbilityScore `bson:"abilityscores"`
}

func (character *Character) SetCharacterName(name string) {
	character.name = name
}

func (character *Character) SetProficiencyBonus(bonus int) {
	character.proficiencyBonus = bonus
}

func (character *Character) AddAbilityScore(name string, value int) {
	var abilityScore AbilityScore
	abilityScore.CreateAbilityScore(name, value)
	character.abilityScores = append(character.abilityScores, abilityScore)
}

func (character *Character) AbilityScoreImprovement(nameInput string, value int) {
	for index, abilityScore := range character.abilityScores {
		if abilityScore.name == nameInput {
			abilityScore.UpdateAbilityScore(value)
			character.abilityScores[index] = abilityScore
			break
		}
	}
}

func (character *Character) PrintCharacterSheet() {
	valueVar := reflect.ValueOf(character).Elem()
	typeVar := valueVar.Type()

	for index := 0; index < valueVar.NumField(); index++ {
		field := typeVar.Field(index)
		value := valueVar.Field(index)

		switch value.Kind() {
		case reflect.String:
			{
				fmt.Printf("%v: %v\n", field.Name, value.String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Int())
			}
		case reflect.Float32, reflect.Float64:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Float())
			}
		case reflect.Bool:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Bool())
			}
		case reflect.Slice:
			{
				if field.Name == "abilityScores" {
					fmt.Println("Ability Scores:")
					for _, abilityScore := range character.abilityScores {
						fmt.Printf("  %s: Score %d, Modifier %d\n", abilityScore.name, abilityScore.score, abilityScore.modifier)
					}
				}
			}
		default:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Interface())
			}
		}
	}
}
