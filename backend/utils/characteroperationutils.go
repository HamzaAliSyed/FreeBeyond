package utils

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IncreaseAbilityScore(character *models.Character, abilityScore string, value int) (*models.Character, error) {
	var updatedIndex int
	var found bool

	for index, score := range character.AbilityScores {
		if abilityScore == score.AbilityName {
			character.AbilityScores[index].Score += value
			character.AbilityScores[index].Modifier = (character.AbilityScores[index].Score - 10) / 2
			updatedIndex = index
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("ability score not found")
	}

	if abilityScore == "Dexterity" {
		var err error
		character, err = SetInitiative(character)
		if err != nil {
			return nil, fmt.Errorf("couldn't update initiative: %w", err)
		}
	}

	for index, save := range character.SavingThrows {
		if save.AttributeName == abilityScore {
			character.SavingThrows[index].Modifer = character.AbilityScores[updatedIndex].Modifier
			character.SavingThrows[index].Value = save.AdditionalBonus + character.SavingThrows[index].Modifer + (int(save.NumberOProficiency) * character.ProficiencyBonus)
		}
	}

	for i, skill := range character.Skills {
		if GetAbilityForSkill(skill.AttributeName) == abilityScore {
			character.Skills[i].Modifer = character.AbilityScores[updatedIndex].Modifier
			character.Skills[i].Value = character.Skills[i].Modifer + character.Skills[i].AdditionalBonus + (int(character.Skills[i].NumberOProficiency) * character.ProficiencyBonus)
		}
	}

	character = CalculatePassives(character)

	return character, nil
}

func GetAbilityForSkill(skillName string) string {
	skillToAbility := map[string]string{
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

	return skillToAbility[skillName]
}

func AddProficiencyToSavingThrow(character *models.Character, ability string, value float64) (*models.Character, error) {
	for index, save := range character.SavingThrows {
		if ability == save.AttributeName {
			character.SavingThrows[index].NumberOProficiency += value
			character.SavingThrows[index].Value = (int(character.SavingThrows[index].NumberOProficiency) * character.ProficiencyBonus) + character.SavingThrows[index].Modifer + character.SavingThrows[index].AdditionalBonus
			return character, nil
		}
	}

	return nil, fmt.Errorf("saving throw not found")
}

func AddProficiencyToSkill(character *models.Character, skillName string, value float64) (*models.Character, error) {
	for index, skill := range character.Skills {
		if skillName == skill.AttributeName {
			character.Skills[index].NumberOProficiency += value
			character.Skills[index].Value = (int(character.Skills[index].NumberOProficiency) * character.ProficiencyBonus) + character.Skills[index].Modifer + character.Skills[index].AdditionalBonus
			if skillName == "Perception" || skillName == "Insight" || skillName == "Investigation" {
				character = CalculatePassives(character)
			}
			return character, nil
		}
	}

	return nil, fmt.Errorf("skill not found")
}

func GiveAdvantageToSavingThrow(character *models.Character, ability string) (*models.Character, error) {
	for index, save := range character.SavingThrows {
		if ability == save.AttributeName {
			character.SavingThrows[index].HasAdvantage = true
			character.SavingThrows[index].HasDisadvantage = false
			return character, nil
		}
	}
	return nil, fmt.Errorf("saving throw not found")
}

func GiveDisadvantageToSavingThrow(character *models.Character, ability string) (*models.Character, error) {
	for index, save := range character.SavingThrows {
		if ability == save.AttributeName {
			character.SavingThrows[index].HasDisadvantage = true
			character.SavingThrows[index].HasAdvantage = false
			return character, nil
		}
	}
	return nil, fmt.Errorf("saving throw not found")
}

func GiveAdvantageToSkill(character *models.Character, skillName string) (*models.Character, error) {
	for index, skill := range character.Skills {
		if skillName == skill.AttributeName {
			character.Skills[index].HasAdvantage = true
			character.Skills[index].HasDisadvantage = false
			return character, nil
		}
	}
	return nil, fmt.Errorf("skill not found")
}

func GiveDisadvantageToSkill(character *models.Character, skillName string) (*models.Character, error) {
	for index, skill := range character.Skills {
		if skillName == skill.AttributeName {
			character.Skills[index].HasDisadvantage = true
			character.Skills[index].HasAdvantage = false
			return character, nil
		}
	}
	return nil, fmt.Errorf("skill not found")
}

func SetInitiative(character *models.Character) (*models.Character, error) {
	for index, dexterity := range character.AbilityScores {
		if dexterity.AbilityName == "Dexterity" {
			character.Initiative.DexterityModifier = character.AbilityScores[index].Modifier
			character.Initiative.Value = character.Initiative.DexterityModifier + character.Initiative.AdditionalBonus
			return character, nil
		}
	}

	return nil, fmt.Errorf("couldnt set initiative")
}

func AddBonusToInitiative(character *models.Character, value int) *models.Character {
	character.Initiative.AdditionalBonus += value
	character.Initiative.Value = character.Initiative.DexterityModifier + character.Initiative.AdditionalBonus
	return character
}

func CalculatePassives(character *models.Character) *models.Character {
	passiveSkills := []string{"Investigation", "Perception", "Insight"}

	for _, passiveSkill := range passiveSkills {
		for _, skill := range character.Skills {
			if skill.AttributeName == passiveSkill {
				switch passiveSkill {
				case "Investigation":
					character.PassiveInvestigation = 10 + skill.Value
				case "Perception":
					character.PassivePerception = 10 + skill.Value
				case "Insight":
					character.PassiveInsight = 10 + skill.Value
				}
				break
			}
		}
	}

	return character
}

func AddBonusToSkill(character *models.Character, skillName string, bonus int) (*models.Character, error) {
	for index, skill := range character.Skills {
		if skillName == skill.AttributeName {
			character.Skills[index].AdditionalBonus += bonus
			character.Skills[index].Value = character.Skills[index].Modifer + character.Skills[index].AdditionalBonus + (int(character.Skills[index].NumberOProficiency) * character.ProficiencyBonus)

			if skillName == "Perception" || skillName == "Insight" || skillName == "Investigation" {
				character = CalculatePassives(character)
			}

			return character, nil
		}
	}

	return nil, fmt.Errorf("skill not found")
}

func AddBonusToSavingThrow(character *models.Character, ability string, bonus int) (*models.Character, error) {
	for index, save := range character.SavingThrows {
		if ability == save.AttributeName {
			character.SavingThrows[index].AdditionalBonus += bonus
			character.SavingThrows[index].Value = character.SavingThrows[index].Modifer + character.SavingThrows[index].AdditionalBonus + (int(character.SavingThrows[index].NumberOProficiency) * character.ProficiencyBonus)
			return character, nil
		}
	}

	return nil, fmt.Errorf("saving throw not found")
}

func SetMovementSpeed(speeds map[string]int, character *models.Character) *models.Character {
	if landSpeed, ok := speeds["land"]; ok {
		character.Movement.LandSpeed = landSpeed
	}
	if swimSpeed, ok := speeds["swim"]; ok {
		character.Movement.SwimSpeed = swimSpeed
	}
	if climbSpeed, ok := speeds["climb"]; ok {
		character.Movement.ClimbSpeed = climbSpeed
	}
	if burrowSpeed, ok := speeds["burrow"]; ok {
		character.Movement.BurrowSpeed = burrowSpeed
	}
	if flySpeed, ok := speeds["fly"]; ok {
		character.Movement.FlySpeed = flySpeed
	}
	return character
}

func AddMovementSpeed(speedType string, value int, character *models.Character) (*models.Character, error) {
	switch speedType {
	case "land":
		character.Movement.LandSpeed += value
	case "swim":
		character.Movement.SwimSpeed += value
	case "climb":
		character.Movement.ClimbSpeed += value
	case "burrow":
		character.Movement.BurrowSpeed += value
	case "fly":
		character.Movement.FlySpeed += value
	default:
		return nil, fmt.Errorf("invalid speed category: %s", speedType)
	}

	return character, nil
}

func GenerateInitialSenses(character *models.Character) *models.Character {
	character.Senses = []models.Sense{
		{SenseName: "Darkvision", SenseRange: 0},
		{SenseName: "Blindsight", SenseRange: 0},
		{SenseName: "Truesight", SenseRange: 0},
		{SenseName: "Tremorsense", SenseRange: 0},
	}
	return character
}

func AddSense(character *models.Character, senseName string, value int) (*models.Character, error) {
	validSenses := map[string]bool{
		"Darkvision":  true,
		"Blindsight":  true,
		"Truesight":   true,
		"Tremorsense": true,
	}

	if !validSenses[senseName] {
		return nil, fmt.Errorf("invalid sense: %s", senseName)
	}

	for i, sense := range character.Senses {
		if sense.SenseName == senseName {
			character.Senses[i].SenseRange = value
			return character, nil
		}
	}

	character.Senses = append(character.Senses, models.Sense{SenseName: senseName, SenseRange: value})
	return character, nil
}

func InitializeEquippedItems(character *models.Character) *models.Character {
	character.EquippedItems = models.EquippedItems{
		Head:    "",
		Torso:   "",
		Legs:    "",
		OffHand: "",
		Hand:    "",
		Wrist:   "",
		SideArm: "",
		Neck:    "",
	}
	return character
}

func EquipItem(character *models.Character, itemName, slotName string) (*models.Character, error) {
	var item models.Items

	if databaseFetchError := database.Items.FindOne(context.TODO(), bson.M{"name": itemName}).Decode(&item); databaseFetchError != nil {
		if databaseFetchError == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("item %s not found in database", itemName)
		}
		return nil, fmt.Errorf("database error: %v", databaseFetchError)
	}

	if item.RequiresAttunement && character.Attunement.NumberOfAttunements >= character.Attunement.MaxAttunement {
		return nil, fmt.Errorf("max attunement reached")
	}

	switch slotName {
	case "Head":
		character.EquippedItems.Head = itemName
	case "Torso":
		character.EquippedItems.Torso = itemName
	case "Legs":
		character.EquippedItems.Legs = itemName
	case "OffHand":
		character.EquippedItems.OffHand = itemName
	case "Hand":
		character.EquippedItems.Hand = itemName
	case "Wrist":
		character.EquippedItems.Wrist = itemName
	case "SideArm":
		character.EquippedItems.SideArm = itemName
	case "Neck":
		character.EquippedItems.Neck = itemName
	default:
		return nil, fmt.Errorf("invalid slot name: %s", slotName)
	}

	return character, nil
}
