package character

import (
	"backend/utils"
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
	skills           []Skill        `bson:"skills"`
	allignment       string         `bson:"allignment"`
}

func CreateCharacter(name, allignment string, strengthscore, dexterityscore, constitutionscore, intelligencescore, wisdomscore, charismascore int) (*Character, error) {
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

	defaultAbilityScore := map[string]int{"Strength": strengthscore,
		"Dexterity":    dexterityscore,
		"Constitution": constitutionscore,
		"Intelligence": intelligencescore,
		"Wisdom":       wisdomscore,
		"Charisma":     charismascore}

	defaultSkills := map[string]string{"Acrobatics": "Dexterity",
		"Animal Handling": "Wisdom",
		"Arcana":          "Intelligence",
		"Athletics":       "Strength",
		"Deception":       "Charisma",
		"History":         "Intelligence",
		"Insight":         "Wisdom",
		"Intimidation":    "Charisma",
		"Investigation":   "Intelligence",
		"Medicine":        "Wisdom",
		"Nature":          "Intelligence",
		"Perception":      "Wisdom",
		"Performance":     "Charisma",
		"Persuasion":      "Charisma",
		"Religion":        "Intelligence",
		"Sleight of Hand": "Dexterity",
		"Stealth":         "Dexterity",
		"Survival":        "Wisdom",
	}

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

	for skillName, abilityName := range defaultSkills {
		var skill Skill
		skill.CreateSkill(skillName, abilityName, character)
		character.skills = append(character.skills, skill)
	}

	validAllignment := []string{"Lawful Good", "Neutral Good", "Chaotic Good", "Lawful Neutral", "True Neutral", "Chaotic Neutral", "Lawful Evil", "Neutral Evil", "Chaotic Evil"}
	isValidAllignment := false
	for _, allignmentValid := range validAllignment {
		if allignmentValid == allignment {
			isValidAllignment = true
		}
	}

	if !isValidAllignment {
		return nil, errors.New("allignment entered is not valid")
	}

	return &character, nil
}

func (character Character) PrintCharacterSheet() {
	fmt.Println("Character Sheet")
	fmt.Printf("Name: %s\n", character.name)
	fmt.Printf("Proficiency Bonus: %v\n", character.proficiencybonus)
	fmt.Printf("Allignment: %s\n", character.allignment)
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

	fmt.Println("Saving Throws")
	for _, savingThrow := range character.savingThrows {
		savingThrow.Print()
	}

	fmt.Println("Skiils")
	for _, skill := range character.skills {
		skill.Print()
	}
}

func (character *Character) CharacterAbilityScore(abilityName string) int {
	var score int
	for _, abilityScore := range character.abilityScores {
		if abilityScore.name == abilityName {
			score = abilityScore.abilityScoreModifier
			break
		}
	}

	return score
}

func (character *Character) RollASkill(skillName string) error {
	fmt.Printf("\nRolling %s for %s\n", skillName, character.name)
	var skillValue int
	var found bool = false
	for _, skill := range character.skills {
		if skill.name == skillName {
			skillValue = skill.value
			found = true
			break
		}
	}

	if !found {
		return errors.New("invalid Skill")
	}

	roller := utils.DieRoller(1, 20)
	rolledValue := roller[0] + skillValue
	fmt.Printf("\n%s rolled %d on %s skill check", character.name, rolledValue, skillName)
	return nil

}
