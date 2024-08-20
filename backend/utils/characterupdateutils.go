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

func AddProfiencyToSkill(character *models.Character, SkillName string) models.Skills {
	skills := character.Skills

	for i := range skills.SkillList {
		if SkillName == skills.SkillList[i].Name {
			fmt.Printf("Before Update, for skill %v final value was %v\n", skills.SkillList[i].Name, skills.SkillList[i].FinalSkillValue)
			skills.SkillList[i].NumberOfProficiencies += 1
			skills.SkillList[i].ProficiencyBonus = character.ProficiencyBonus
			skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue +
				(int(skills.SkillList[i].NumberOfProficiencies) * character.ProficiencyBonus) +
				skills.SkillList[i].AdditionalBoostValue
			fmt.Printf("Before Update, for skill %v final value was %v\n", skills.SkillList[i].Name, skills.SkillList[i].FinalSkillValue)
			break
		}
	}

	return skills
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

func UpdateSavingThrowsAfterASI(character *models.Character) []models.SavingThrow {
	for i := range character.SavingThrow {
		switch character.SavingThrow[i].Attribute {
		case "Strength":
			if character.SavingThrow[i].AttributeModifier != character.Modifiers.StrengthModifier {
				character.SavingThrow[i].AttributeModifier = character.Modifiers.StrengthModifier
				character.SavingThrow[i].SavingThrowValue = character.SavingThrow[i].AttributeModifier + (character.SavingThrow[i].NumberOfProficiencies * character.ProficiencyBonus)
				fmt.Printf("Saving throw of Strength is now %v \n", character.SavingThrow[i].SavingThrowValue)
			}
		case "Dexterity":
			if character.SavingThrow[i].AttributeModifier != character.Modifiers.DexterityModifier {
				character.SavingThrow[i].AttributeModifier = character.Modifiers.DexterityModifier
				character.SavingThrow[i].SavingThrowValue = character.SavingThrow[i].AttributeModifier + (character.SavingThrow[i].NumberOfProficiencies * character.ProficiencyBonus)
				fmt.Printf("Saving throw of Dexterity is now %v\n", character.SavingThrow[i].SavingThrowValue)
			}
		case "Constitution":
			if character.SavingThrow[i].AttributeModifier != character.Modifiers.ConstitutionModifier {
				character.SavingThrow[i].AttributeModifier = character.Modifiers.ConstitutionModifier
				character.SavingThrow[i].SavingThrowValue = character.SavingThrow[i].AttributeModifier + (character.SavingThrow[i].NumberOfProficiencies * character.ProficiencyBonus)
				fmt.Printf("Saving throw of Constitution is now %v\n", character.SavingThrow[i].SavingThrowValue)
			}
		case "Intelligence":
			if character.SavingThrow[i].AttributeModifier != character.Modifiers.IntelligenceModifier {
				character.SavingThrow[i].AttributeModifier = character.Modifiers.IntelligenceModifier
				character.SavingThrow[i].SavingThrowValue = character.SavingThrow[i].AttributeModifier + (character.SavingThrow[i].NumberOfProficiencies * character.ProficiencyBonus)
				fmt.Printf("Saving throw of Intelligence is now %v\n", character.SavingThrow[i].SavingThrowValue)
			}
		case "Wisdom":
			if character.SavingThrow[i].AttributeModifier != character.Modifiers.WisdomModifier {
				character.SavingThrow[i].AttributeModifier = character.Modifiers.WisdomModifier
				character.SavingThrow[i].SavingThrowValue = character.SavingThrow[i].AttributeModifier + (character.SavingThrow[i].NumberOfProficiencies * character.ProficiencyBonus)
				fmt.Printf("Saving throw of Wisdom is now %v\n", character.SavingThrow[i].SavingThrowValue)
			}
		case "Charisma":
			if character.SavingThrow[i].AttributeModifier != character.Modifiers.CharismaModifier {
				character.SavingThrow[i].AttributeModifier = character.Modifiers.CharismaModifier
				character.SavingThrow[i].SavingThrowValue = character.SavingThrow[i].AttributeModifier + (character.SavingThrow[i].NumberOfProficiencies * character.ProficiencyBonus)
				fmt.Printf("Saving throw of Charisma is now %v\n", character.SavingThrow[i].SavingThrowValue)
			}
		}
	}

	return character.SavingThrow
}

func UpdateSkillsAfterASI(character *models.Character, ability string) models.Skills {
	var skills = character.Skills
	for i := range skills.SkillList {
		if skills.SkillList[i].AssociatedAttribute == ability {
			switch ability {
			case "Strength":
				skills.SkillList[i].AssociatedAttributeValue = character.Modifiers.StrengthModifier
				skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue + (skills.SkillList[i].ProficiencyBonus * int(skills.SkillList[i].NumberOfProficiencies)) + skills.SkillList[i].AdditionalBoostValue
				fmt.Printf("For skill %v attributed value %v was increased to %v\n", skills.SkillList[i].Name, skills.SkillList[i].AssociatedAttribute, skills.SkillList[i].AssociatedAttributeValue)
			case "Dexterity":
				skills.SkillList[i].AssociatedAttributeValue = character.Modifiers.DexterityModifier
				skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue + (skills.SkillList[i].ProficiencyBonus * int(skills.SkillList[i].NumberOfProficiencies)) + skills.SkillList[i].AdditionalBoostValue
				fmt.Printf("For skill %v attributed value %v was increased to %v\n", skills.SkillList[i].Name, skills.SkillList[i].AssociatedAttribute, skills.SkillList[i].AssociatedAttributeValue)
			case "Constitution":
				skills.SkillList[i].AssociatedAttributeValue = character.Modifiers.ConstitutionModifier
				skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue + (skills.SkillList[i].ProficiencyBonus * int(skills.SkillList[i].NumberOfProficiencies)) + skills.SkillList[i].AdditionalBoostValue
				fmt.Printf("For skill %v attributed value %v was increased to %v\n", skills.SkillList[i].Name, skills.SkillList[i].AssociatedAttribute, skills.SkillList[i].AssociatedAttributeValue)
			case "Intelligence":
				skills.SkillList[i].AssociatedAttributeValue = character.Modifiers.IntelligenceModifier
				skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue + (skills.SkillList[i].ProficiencyBonus * int(skills.SkillList[i].NumberOfProficiencies)) + skills.SkillList[i].AdditionalBoostValue
				fmt.Printf("For skill %v attributed value %v was increased to %v\n", skills.SkillList[i].Name, skills.SkillList[i].AssociatedAttribute, skills.SkillList[i].AssociatedAttributeValue)
			case "Wisdom":
				skills.SkillList[i].AssociatedAttributeValue = character.Modifiers.WisdomModifier
				skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue + (skills.SkillList[i].ProficiencyBonus * int(skills.SkillList[i].NumberOfProficiencies)) + skills.SkillList[i].AdditionalBoostValue
				fmt.Printf("For skill %v attributed value %v was increased to %v\n", skills.SkillList[i].Name, skills.SkillList[i].AssociatedAttribute, skills.SkillList[i].AssociatedAttributeValue)
			case "Charisma":
				skills.SkillList[i].AssociatedAttributeValue = character.Modifiers.CharismaModifier
				skills.SkillList[i].FinalSkillValue = skills.SkillList[i].AssociatedAttributeValue + (skills.SkillList[i].ProficiencyBonus * int(skills.SkillList[i].NumberOfProficiencies)) + skills.SkillList[i].AdditionalBoostValue
				fmt.Printf("For skill %v attributed value %v was increased to %v\n", skills.SkillList[i].Name, skills.SkillList[i].AssociatedAttribute, skills.SkillList[i].AssociatedAttributeValue)
			default:
				fmt.Printf("Unrecognized ability: %v\n", ability)
			}
		}
	}

	return skills
}

func UpdateMaxCarryWeight(character *models.Character) int {
	strengthScore := character.MainAttributes.StrengthScore
	carryweightmax := strengthScore * 15
	return carryweightmax
}

func AddAdvantageToSkill(character *models.Character, SkillName string) models.Skills {
	skills := character.Skills

	for i := range skills.SkillList {
		if SkillName == skills.SkillList[i].Name {
			skills.SkillList[i].HasAdvantage = true
			break
		}
	}

	return skills
}

func AddAdvantageToSavingThrows(character *models.Character, SavingThrow string) []models.SavingThrow {
	SavingThrows := character.SavingThrow

	for i := range SavingThrows {
		if SavingThrows[i].Attribute == SavingThrow {
			SavingThrows[i].HasAdvantage = true
			break
		}
	}

	return SavingThrows
}

func AddProfiencyToSavingThrow(character *models.Character, SavingThrow string) []models.SavingThrow {
	savingthrows := character.SavingThrow

	for i := range savingthrows {
		if SavingThrow == savingthrows[i].Attribute {
			savingthrows[i].NumberOfProficiencies += 1
			savingthrows[i].SavingThrowValue = savingthrows[i].AttributeModifier + (savingthrows[i].NumberOfProficiencies * character.ProficiencyBonus)
			break
		}
	}

	return savingthrows
}

func GenerateGenericMeleeAttack(character *models.Character) models.AnAttack {
	var attack models.AnAttack
	damage := map[string]string{
		"Bludgeoning": "1d4+str",
	}

	attack.Name = "Unarmed Attack"
	attack.Type = "Melee"
	attack.Range = 5
	attack.RangeMax = 5
	attack.RangeMin = 0
	attack.AttributeUsed = "strength"
	attack.AttributeValue = character.Modifiers.StrengthModifier
	attack.Damage = damage

	return attack
}

func CharacterInitiative(character *models.Character) int {
	initiative := character.Modifiers.DexterityModifier
	return initiative
}

func GenerateFeatureForLevel(featurename []string, featuretype []string, featureresetinformation []string) []models.ClassFeature {
	features := []models.ClassFeature{}

	for i := 0; i < len(featurename); i++ {
		var feature models.ClassFeature
		feature.Name = featurename[i]
		feature.Type = featuretype[i]
		feature.ResetInformation = featureresetinformation[i]
		features = append(features, feature)
	}

	return features
}
