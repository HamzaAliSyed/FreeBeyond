package classes

import (
	"backend/models"
	"backend/utils"
	"fmt"
)

func BloodHunterLevelOne(character *models.Character) (*models.Character, error) {

	intelligenceCheck, intelligenceCheckError := utils.GetAbilityScore(character, "Intelligence")
	if intelligenceCheckError != nil {
		return nil, intelligenceCheckError
	}
	if intelligenceCheck <= 13 {
		return nil, fmt.Errorf("intelligence must be greater than 13 for Blood Hunter")
	}

	strengthScore, strengthErr := utils.GetAbilityScore(character, "Strength")
	if strengthErr != nil {
		return nil, strengthErr
	}

	dexterityScore, dexterityErr := utils.GetAbilityScore(character, "Dexterity")
	if dexterityErr != nil {
		return nil, dexterityErr
	}

	if strengthScore <= 13 && dexterityScore <= 13 {
		return nil, fmt.Errorf("either Strength or Dexterity must be greater than 13 for Blood Hunter")
	}

	if character.Class.ClassesAndLevels == nil {
		character.Class.ClassesAndLevels = make(map[string]int)
	}

	character.Class.ClassesAndLevels["Bloodhunter"] = 1

	if character.Class.HitDices == nil {
		character.Class.HitDices = make(map[models.HitDie]int)
	}
	character.Class.HitDices["d10"] += 1

	character = utils.AddWeaponProficiencies(character, []string{"Simple Weapon", "Martial Weapons"})
	character = utils.AddArmorProficiencies(character, []string{"Light Armor", "Medium Armor", "Shields"})
	character = utils.AddToolProficiencies(character, []string{"Alchemist's supplies"})
	character, _ = utils.AddProficiencyToSavingThrow(character, "Dexterity", 1.0)
	character, _ = utils.AddProficiencyToSavingThrow(character, "Intelligence", 1.0)

	return character, nil

}
