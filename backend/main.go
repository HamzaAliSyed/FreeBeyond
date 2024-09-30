package main

import (
	character "backend/Character"
	"fmt"
	"log"
)

func main() {
	Thorgar, err := character.CreateCharacter("Thorgar Ragehammer", "Lawful Neutral", 18, 15, 18, 10, 12, 14)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Exiting Program")
		log.Fatal("Invalid Character")
	} else {
		Thorgar.PrintCharacterSheet()
	}

	skillRollError := Thorgar.RollASkill("Persuasion")
	if skillRollError != nil {
		fmt.Println(skillRollError.Error())
	}
}
