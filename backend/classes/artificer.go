package classes

import "backend/models"

func ArtificerLevelOne(character *models.Character, choices map[string]string) {
	if character.HitDie == "" {
		character.HitDie = "d8"
	} else {
		hitdie := models.CompareHitDie(character.HitDie, "d8")
		character.HitDie = hitdie
	}

	character.CalculateHitPoints(0)
	character.AddArmorProficiency("light armor", "medium armor", "shields")

	if choices["Optional Rule"] == "yes" {
		character.AddWeaponProficiency("simple weapons", "firearms")
	} else {
		character.AddWeaponProficiency("simple weapons")
	}

	character.AddToolProficiencies("thieves' tools")
	character.AddToolProficiencies("tinker's tools")
	character.AddToolProficiencies(choices["third tool proficiency"])
	character.AddClassFeatures("Magical Tinkering")
	character.CanDoSpellCasting = true
}
