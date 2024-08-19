package utils

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var MainAttributesInstance models.MainAttributes
var ModiferInstance models.Modifiers

func UpdateAbilityScore(characterID string, abilityscore string, value int) (models.MainAttributes, models.Modifiers) {
	character, characterretrieveErr := RetrieveCharacter(characterID, database.Characters)
	if characterretrieveErr != nil {
		fmt.Println("Could not find the character from ID")
		log.Fatal("Invalid ID")
	}
	MainAttributesInstance = character.MainAttributes
	ModiferInstance = character.Modifiers
	switch abilityscore {
	case "strength":
		{
			MainAttributesInstance.StrengthScore += value
			ModiferInstance.StrengthModifier = ModifierCalculator(MainAttributesInstance.StrengthScore)
		}
	case "dexterity":
		MainAttributesInstance.DexterityScore += value
		ModiferInstance.DexterityModifier = ModifierCalculator(MainAttributesInstance.DexterityScore)
	case "constitution":
		MainAttributesInstance.ConstitutionScore += value
		ModiferInstance.ConstitutionModifier = ModifierCalculator(MainAttributesInstance.ConstitutionScore)
	case "intelligence":
		MainAttributesInstance.IntelligenceScore += value
		ModiferInstance.IntelligenceModifier = ModifierCalculator(MainAttributesInstance.IntelligenceScore)
	case "wisdom":
		MainAttributesInstance.WisdomScore += value
		ModiferInstance.WisdomModifier = ModifierCalculator(MainAttributesInstance.WisdomScore)
	case "charisma":
		MainAttributesInstance.CharismaScore += value
		ModiferInstance.CharismaModifier = ModifierCalculator(MainAttributesInstance.CharismaScore)
	default:
		// Handle invalid ability score
		fmt.Println(fmt.Errorf("invalid ability score: %s", abilityscore))
	}

	return MainAttributesInstance, ModiferInstance

}

func AddProfiencyToSkill(characterID string, SkillName string) {
	character, characterretrieveErr := RetrieveCharacter(characterID, database.Characters)
	if characterretrieveErr != nil {
		fmt.Println("Could not find the character from ID")
		log.Fatal("Invalid ID")
	}
	for _, skill := range character.Skills.SkillList {
		if skill.Name == SkillName {
			toadd := GetProficiencyBonus(characterID)
			skill.NumberOfProficiencies += 1
			skill.ProficiencyBonus = toadd
			skill.FinalSkillValue = skill.AssociatedAttributeValue + (skill.ProficiencyBonus * int(skill.NumberOfProficiencies)) + skill.AdditionalBoostValue
			fmt.Printf("Updated %s to have proficiency, now final value is %v", skill.Name, skill.FinalSkillValue)
			break
		}
	}
	update := bson.D{{Key: "$set", Value: character}}
	_, err := database.Characters.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: character.ID}}, update)
	if err != nil {
		fmt.Println(fmt.Errorf("could not update character: %w", err))
	}
}

func UpdateCharacterToDB(character *models.Character) {
	filter := bson.D{{Key: "_id", Value: character.ID}}
	update := bson.D{{Key: "$set", Value: character}}
	_, updateerr := database.Characters.UpdateOne(context.TODO(), filter, update)
	if updateerr != nil {
		fmt.Println(fmt.Errorf("could not update character: %v", updateerr))
	} else {
		println("Update of Character was successful")
	}
}

func UpdateSavingThrowsAfterASI(character *models.Character) *[]models.SavingThrow {
	for _, savingThrow := range character.SavingThrow {
		switch savingThrow.Attribute {
		case "Strength":
			{
				if savingThrow.AttributeModifier != character.Modifiers.StrengthModifier {
					savingThrow.AttributeModifier = character.Modifiers.StrengthModifier
					savingThrow.SavingThrowValue = savingThrow.AttributeModifier + (savingThrow.NumberOfProficiencies * character.ProficiencyBonus)
					fmt.Printf("Saving throw of Strength is now %v", savingThrow.SavingThrowValue)
				}
			}
		case "Dexterity":
			{
				if savingThrow.AttributeModifier != character.Modifiers.DexterityModifier {
					savingThrow.AttributeModifier = character.Modifiers.DexterityModifier
					savingThrow.SavingThrowValue = savingThrow.AttributeModifier + (savingThrow.NumberOfProficiencies * character.ProficiencyBonus)
					fmt.Printf("Saving throw of Dexterity is now %v", savingThrow.SavingThrowValue)
				}
			}
		case "Constitution":
			{
				if savingThrow.AttributeModifier != character.Modifiers.ConstitutionModifier {
					savingThrow.AttributeModifier = character.Modifiers.ConstitutionModifier
					savingThrow.SavingThrowValue = savingThrow.AttributeModifier + (savingThrow.NumberOfProficiencies * character.ProficiencyBonus)
					fmt.Printf("Saving throw of Constitution is now %v", savingThrow.SavingThrowValue)
				}
			}
		case "Intelligence":
			{
				if savingThrow.AttributeModifier != character.Modifiers.IntelligenceModifier {
					savingThrow.AttributeModifier = character.Modifiers.IntelligenceModifier
					savingThrow.SavingThrowValue = savingThrow.AttributeModifier + (savingThrow.NumberOfProficiencies * character.ProficiencyBonus)
					fmt.Printf("Saving throw of Intelligence is now %v", savingThrow.SavingThrowValue)
				}
			}
		case "Wisdom":
			{
				if savingThrow.AttributeModifier != character.Modifiers.WisdomModifier {
					savingThrow.AttributeModifier = character.Modifiers.WisdomModifier
					savingThrow.SavingThrowValue = savingThrow.AttributeModifier + (savingThrow.NumberOfProficiencies * character.ProficiencyBonus)
					fmt.Printf("Saving throw of Wisdom is now %v", savingThrow.SavingThrowValue)
				}
			}
		case "Charisma":
			{
				if savingThrow.AttributeModifier != character.Modifiers.CharismaModifier {
					savingThrow.AttributeModifier = character.Modifiers.CharismaModifier
					savingThrow.SavingThrowValue = savingThrow.AttributeModifier + (savingThrow.NumberOfProficiencies * character.ProficiencyBonus)
					fmt.Printf("Saving throw of Charisma is now %v", savingThrow.SavingThrowValue)
				}
			}
		}
	}

	return &character.SavingThrow
}
