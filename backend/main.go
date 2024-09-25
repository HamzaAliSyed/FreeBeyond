package main

import character "backend/Character"

func main() {
	ThorgarRagehammer := character.CreateCharacter("Thorgar Ragehammer")
	ThorgarRagehammer.PrintCharacterSheet()
}
