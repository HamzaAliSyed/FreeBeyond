package character

import (
	"errors"
	"fmt"
)

type HitDie string

const (
	d4  = "d4"
	d6  = "d6"
	d8  = "d8"
	d10 = "d10"
	d12 = "d12"
	d20 = "d20"
)

type Character struct {
	name             string         `bson:"name"`
	proficiencybonus int            `bson:"proficiencybonus"`
	hitdie           map[HitDie]int `bson:"hitdie"`
	abilityScores    []AbilityScore `bson:"abilityscores"`
	savingThrows     []SavingThrow  `bson:"savingthrow"`
}

func CreateCharacter(name string, strengthscore, dexterityscore, constitutionscore, intelligencescore, wisdomscore, charismascore int) (*Character, error) {
	var character Character
	character.name = name
	character.proficiencybonus = 2
	character.hitdie = map[HitDie]int{
		d4:  0,
		d6:  0,
		d8:  0,
		d10: 0,
		d12: 0,
		d20: 0,
	}

	defaultAbilityScore := map[string]int{"Strength": strengthscore, "Dexterity": dexterityscore, "Constitution": constitutionscore, "Intelligence": intelligencescore, "Wisdom": wisdomscore, "Charisma": charismascore}

	for abilityScoreName, AbilityScoreValue := range defaultAbilityScore {
		var abilityScore AbilityScore
		creationError := abilityScore.CreateAbilityScore(abilityScoreName, AbilityScoreValue)
		if creationError != nil {
			return nil, errors.New("couldn`t create character as provided ability scores are invalid")
		}
		character.abilityScores = append(character.abilityScores, abilityScore)
	}

	for _, abilityScore := range character.abilityScores {
		name, mod := abilityScore.GetNameandMod()
		var savingThrow SavingThrow
		savingThrow.CreateSavingThrow(name, mod)
		character.savingThrows = append(character.savingThrows, savingThrow)
	}

	return &character, nil
}

func (character Character) PrintCharacterSheet() {
	fmt.Println("Character Sheet")
	fmt.Printf("Name: %s\n", character.name)
	fmt.Printf("Proficiency Bonus: %v\n", character.proficiencybonus)
	fmt.Printf("HITDIES \n")
	for singlehitdie, hitdievalue := range character.hitdie {
		if hitdievalue > 0 {
			fmt.Printf("%v : %v\n", singlehitdie, hitdievalue)
		}
	}
	fmt.Println("Ability Scores")
	for _, singleAS := range character.abilityScores {
		singleAS.Print()
	}

	fmt.Println("SavingThrows")
	for _, savingThrow := range character.savingThrows {
		savingThrow.Print()
	}
}
