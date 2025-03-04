package models

import "fmt"

type Attack interface {
	Name() string
	Print()
}

type ACBeatingAttack struct {
	name           string
	AttackModifier int
	attackDamages  []AttackDamage
}

func (acBeatingAttack ACBeatingAttack) Name() string {
	return acBeatingAttack.name
}

func (acBeatingAttack ACBeatingAttack) Print() {
	fmt.Println("We are printing Attack")
	fmt.Println("Attack Name: ", acBeatingAttack.name)
	fmt.Println("Attack Modifier: ", acBeatingAttack.AttackModifier)
	fmt.Println("Now Printing Damages")
	for _, attackDamage := range acBeatingAttack.attackDamages {
		attackDamage.Print()
	}
}

type AttackParameters struct {
	Name               string
	MainAbility        string
	Traits             []string
	NumberOfDie        []int
	ArrayOfDamageTypes []string
	ArrayOfHitDies     []string
}

func GenerateACBeatingAttackForCharacter(character *Character, parameters AttackParameters) (*ACBeatingAttack, error) {
	var acBeatingAttack ACBeatingAttack
	acBeatingAttack.name = parameters.Name
	attackDamageArrays, attackTypesError := CreateAttackType(parameters.NumberOfDie, parameters.ArrayOfDamageTypes, parameters.ArrayOfHitDies)
	if attackTypesError != nil {
		return nil, fmt.Errorf("error encoding damage types: \n%v", attackTypesError)
	}
	modifier, modifierError := GenerateAttackModifier(*character, parameters.MainAbility)
	if modifierError != nil {
		return nil, fmt.Errorf("cannot Generate Modifier %v", modifierError)
	}
	acBeatingAttack.AttackModifier = modifier
	acBeatingAttack.attackDamages = attackDamageArrays
	return &acBeatingAttack, nil
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

type HitDie string

const (
	D4  HitDie = "d4"
	D6  HitDie = "d6"
	D8  HitDie = "d8"
	D10 HitDie = "d10"
	D12 HitDie = "d12"
)

type AttackDamage struct {
	numberOfDies int
	damage       DamageTypes
	hitDie       HitDie
}

func (attackDamage AttackDamage) Print() {
	fmt.Printf("%d %s %s\n", attackDamage.numberOfDies, attackDamage.damage, attackDamage.hitDie)
}

func CreateAttackType(numbersOfDices []int, numberOfDamageTypes []string, numberOfHitDices []string) ([]AttackDamage, error) {
	var attackDamages []AttackDamage
	if len(numbersOfDices) != len(numberOfDamageTypes) && len(numbersOfDices) != len(numberOfHitDices) {
		return nil, fmt.Errorf("bad Input Array")
	}

	validDamageTypes := map[DamageTypes]bool{
		Fire:        true,
		Lightning:   true,
		Cold:        true,
		Acid:        true,
		Poison:      true,
		Force:       true,
		Thunder:     true,
		Slashing:    true,
		Piercing:    true,
		Bludgeoning: true,
		Radiant:     true,
		Necrotic:    true,
		Psychic:     true,
	}

	validHitDies := map[HitDie]bool{
		D4:  true,
		D6:  true,
		D8:  true,
		D10: true,
		D12: true,
	}

	for index := 0; index < len(numbersOfDices); index++ {
		damageType := DamageTypes(numberOfDamageTypes[index])
		if !validDamageTypes[damageType] {
			return nil, fmt.Errorf("invalid damage types: %s", numberOfDamageTypes[index])
		}

		hitDie := HitDie(numberOfHitDices[index])
		if !validHitDies[hitDie] {
			return nil, fmt.Errorf("invalid hit die: %s", numberOfHitDices[index])
		}

		attackDamage := AttackDamage{
			numberOfDies: numbersOfDices[index],
			damage:       damageType,
			hitDie:       hitDie,
		}

		attackDamages = append(attackDamages, attackDamage)
	}

	return attackDamages, nil
}

func GenerateAttackModifier(character Character, mainAttribute string) (int, error) {
	abilityScore, abilityScoreExist := character.AbilityScores[mainAttribute]
	if !abilityScoreExist {
		return 0, fmt.Errorf("ability score %s doesnt exist in the character", mainAttribute)
	}

	asStruct, okEncoding := abilityScore.(AbilityScore)
	if !okEncoding {
		return 0, fmt.Errorf("cannot encode the ability struct")
	}

	modifier := asStruct.Modifier
	return modifier, nil
}
