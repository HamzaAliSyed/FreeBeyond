package character

import (
	item "backend/Item"
	"fmt"
	"reflect"
)

type Character struct {
	name             string            `bson:"name"`
	proficiencyBonus int               `bson:"proficiencybonus"`
	abilityScores    []AbilityScore    `bson:"abilityscores"`
	inventory        map[item.Item]int `bson:"inventory"`
	equippedItems    Equipment         `bson:"euippeditems"`
}

func (character *Character) NewCharacter() {
	character.inventory = make(map[item.Item]int)
}

func (character *Character) SetCharacterName(name string) {
	character.name = name
}

func (character *Character) SetProficiencyBonus(bonus int) {
	character.proficiencyBonus = bonus
}

func (character *Character) AddAbilityScore(name string, value int) {
	var abilityScore AbilityScore
	abilityScore.CreateAbilityScore(name, value)
	character.abilityScores = append(character.abilityScores, abilityScore)
}

func (character *Character) AbilityScoreImprovement(nameInput string, value int) {
	for index, abilityScore := range character.abilityScores {
		if abilityScore.name == nameInput {
			abilityScore.UpdateAbilityScore(value)
			character.abilityScores[index] = abilityScore
			break
		}
	}
}

func (character *Character) AddItemToInventory(item item.Item) {

	if count, exists := character.inventory[item]; exists {
		character.inventory[item] = count + 1
	} else {
		character.inventory[item] = 1
	}
}

func (character *Character) RemoveOneItemFromInventory(item item.Item) {
	if count, exists := character.inventory[item]; exists {
		if count > 1 {
			character.inventory[item] = count - 1
		} else {
			delete(character.inventory, item)
		}
	}
}

func (character *Character) EquipWeapon(item item.Item) {
	weaponProperties := item.GetItemTags()
	for _, weaponProperty := range weaponProperties {
		switch weaponProperty {
		case "Versatile":
			{
				if character.equippedItems.mainhand == "" && character.equippedItems.offhand == "" {
					character.equippedItems.mainhand = item.GetItemName()
					character.equippedItems.offhand = item.GetItemName()
				} else {
					character.equippedItems.mainhand = item.GetItemName()
				}
			}
		case "Two Handed":
			{
				character.equippedItems.mainhand = item.GetItemName()
				character.equippedItems.offhand = item.GetItemName()
			}
		case "Light":
			{
				if character.equippedItems.mainhand == "" {
					character.equippedItems.mainhand = item.GetItemName()
				} else {
					character.equippedItems.offhand = item.GetItemName()
				}
			}
		default:
			{
				character.equippedItems.mainhand = item.GetItemName()
			}
		}
	}
}

func (character *Character) UnequipAnItem(item item.Item) {
	itemName := item.GetItemName()
	value := reflect.ValueOf(character.equippedItems).Elem()

	for index := 0; index < value.NumField(); index++ {
		field := value.Field(index)
		if field.String() == itemName {
			field.SetString("")
		}
	}
}

func (character *Character) PrintCharacterSheet() {
	valueVar := reflect.ValueOf(character).Elem()
	typeVar := valueVar.Type()

	for index := 0; index < valueVar.NumField(); index++ {
		field := typeVar.Field(index)
		value := valueVar.Field(index)

		switch value.Kind() {
		case reflect.String:
			{
				fmt.Printf("%v: %v\n", field.Name, value.String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Int())
			}
		case reflect.Float32, reflect.Float64:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Float())
			}
		case reflect.Bool:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Bool())
			}
		case reflect.Slice:
			{
				if field.Name == "abilityScores" {
					fmt.Println("Ability Scores:")
					for _, abilityScore := range character.abilityScores {
						fmt.Printf("  %s: Score %d, Modifier %d\n", abilityScore.name, abilityScore.score, abilityScore.modifier)
					}
				}
			}
		case reflect.Map:
			{
				if field.Name == "inventory" {
					fmt.Println("Inventory:")
					for item, count := range character.inventory {
						fmt.Printf("  %s: %d\n", item.GetItemName(), count)
					}
				}
			}
		case reflect.Struct:
			if field.Name == "equippedItems" {
				fmt.Println("Equipped Items:")
				equipmentValue := reflect.ValueOf(character.equippedItems)
				equipmentType := equipmentValue.Type()
				for i := 0; i < equipmentValue.NumField(); i++ {
					equipField := equipmentType.Field(i)
					equipValue := equipmentValue.Field(i)
					if equipValue.String() != "" {
						fmt.Printf("  %s: %s\n", equipField.Name, equipValue.String())
					}
				}
			}
		default:
			{
				fmt.Printf("%v: %v\n", field.Name, value.Interface())
			}
		}
	}
}
