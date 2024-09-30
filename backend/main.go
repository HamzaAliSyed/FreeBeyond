package main

import (
	character "backend/Character"
	"fmt"
)

func main() {
	Thorgar, err := character.CreateCharacter("Thorgar Ragehammer", 18, 15, 18, 10, 12, 14)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		Thorgar.PrintCharacterSheet()
	}

	Thorgar.RollASkill("Intimidation")

}
