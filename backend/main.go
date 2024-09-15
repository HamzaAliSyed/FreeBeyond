package main

import (
	"backend/models"
	"fmt"
)

/*import (
	"backend/database"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	environmentPath := "../.env"
	environmenterror := godotenv.Load(environmentPath)
	if environmenterror != nil {
		log.Fatalf("Error loading .env file")
	}

	const port = "2712"
	database.ConnectToMongo()
	backend := http.NewServeMux()

	log.Fatal(http.ListenAndServe(":"+port, backend))
}*/

func main() {
	// Create a new character
	var Thorgar models.Character
	Thorgar.SetCharacterName("Thorgar Ragehammer")
	Thorgar.SetProficiencyBonus(2)

	abilityScoreNames := [6]string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}
	abilityScoreValues := [6]int{18, 15, 14, 10, 12, 14}

	for index, abilityScore := range abilityScoreNames {
		Thorgar.SetCharacterAbilityScore(abilityScore, abilityScoreValues[index])
		Thorgar.CreateCharacterSavingThrow(abilityScore)
	}

	fmt.Println("Character Name:", Thorgar.GetCharacterName())
	fmt.Println("Proficiency Bonus:", Thorgar.GetProficiencyBonus())
	fmt.Println("\nAbility Scores:")
	for _, abilityScore := range Thorgar.GetAllCharacterAbilityScores() {
		fmt.Printf(" - %s: %d (Modifier: %d)\n", abilityScore.GetAbilityName(), abilityScore.GetAbilityScore(), abilityScore.GetAbilityModifier())
	}
	fmt.Println("\nSaving Throws:")
	for _, savingThrow := range Thorgar.GetCharacterSavingThrows() {
		fmt.Printf(" - %s: %d\n", savingThrow.GetSavingThrowName(), savingThrow.GetSavingThrowValue())
	}

	fmt.Println("\n Now Buffing Thorgar more")
	Thorgar.UpdateAbilityScore("Strength", 2)
	fmt.Println("Adding 2 to strength")
	fmt.Println("Character Name:", Thorgar.GetCharacterName())
	fmt.Println("Proficiency Bonus:", Thorgar.GetProficiencyBonus())
	fmt.Println("\nAbility Scores:")
	for _, abilityScore := range Thorgar.GetAllCharacterAbilityScores() {
		fmt.Printf(" - %s: %d (Modifier: %d)\n", abilityScore.GetAbilityName(), abilityScore.GetAbilityScore(), abilityScore.GetAbilityModifier())
	}
	fmt.Println("\nSaving Throws:")
	for _, savingThrow := range Thorgar.GetCharacterSavingThrows() {
		fmt.Printf(" - %s: %d\n", savingThrow.GetSavingThrowName(), savingThrow.GetSavingThrowValue())
	}
	skillMap := map[string]string{
		"Acrobatics":      "Dexterity",
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
	for skillName, attribute := range skillMap {
		Thorgar.CreateNewSkill(skillName, attribute)
	}
	fmt.Println("\nSkills:")
	for _, skill := range Thorgar.GetAllSkills() {
		fmt.Printf(" - %s (%s): %d\n", skill.GetSkillName(), skill.GetSkillAttribute(), skill.GetValue())
	}
	fmt.Println("\nWe are now going to give proficiency to Athletics and Nature")
	Thorgar.AddProficiencyToSkill("Athletics", 1)
	Thorgar.AddProficiencyToSkill("Nature", 0.5)
	fmt.Println("\nUpdated Skills:")
	for _, skill := range Thorgar.GetAllSkills() {
		fmt.Printf(" - %s (%s): %d\n", skill.GetSkillName(), skill.GetSkillAttribute(), skill.GetValue())
	}
}
