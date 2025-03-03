package models

type Attack interface {
	Name() string
}

type ACBeatingAttack struct {
	name string
}

func (acBeatingAttack ACBeatingAttack) Name() string {
	return acBeatingAttack.name
}

type AttackParameters struct {
	MainAbility string
	Traits      []string
}

func GenerateACBeatingAttackForCharacter(character *Character) (ACBeatingAttack, error) {
	var acBeatingAttack ACBeatingAttack
	return acBeatingAttack, nil
}

type DamageTypes string

const (
	Fire        DamageTypes = "Fire"
	Lightning   DamageTypes = "Lightning"
	Cold        DamageTypes = "Cold"
	Acid        DamageTypes = "Acid"
	Poison      DamageTypes = "Poison"
	Force       DamageTypes = "Force"
	Thunder     DamageTypes = "Thunder"
	Slashing    DamageTypes = "Slashing"
	Piercing    DamageTypes = "Piercing"
	Bludgeoning DamageTypes = "Bludgeoning"
	Radiant     DamageTypes = "Radiant"
	Necrotic    DamageTypes = "Necrotic"
	Psychic     DamageTypes = "Psychic"
)
