package character

import "fmt"

type Character struct {
	name             string `bson:"name"`
	proficiencybonus int    `bson:"proficiencybonus"`
}

func CreateCharacter(name string) Character {
	var character Character
	character.name = name
	character.proficiencybonus = 2

	return character
}

func (character Character) PrintCharacterSheet() {
	fmt.Println("Character Sheet")
	fmt.Printf("Name: %s\n", character.name)
	fmt.Printf("Proficiency Bonus: %s\n", character.proficiencybonus)
}
