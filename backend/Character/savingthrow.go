package character

import "fmt"

type SavingThrow struct {
	name                  string  `bson:"name"`
	scoreModifier         int     `bson:"scoremodifier"`
	additionalBonus       int     `bson:"additionalbonus"`
	hasAdvantage          bool    `bson:"hasadvantage"`
	hasDisadvantage       bool    `bson:"hasDisadvantage"`
	numberOfProficiencies float64 `bson:"numberofproficiencies"`
	value                 int     `bson:"value"`
}

func (savingThrow *SavingThrow) CreateSavingThrow(name string, modifier int) {
	savingThrow.name = name
	savingThrow.scoreModifier = modifier
	savingThrow.additionalBonus = 0
	savingThrow.hasAdvantage = false
	savingThrow.hasDisadvantage = false
	savingThrow.numberOfProficiencies = 0
	savingThrow.value = modifier

}

func (savingThrow *SavingThrow) Print() {
	fmt.Printf("\nSaving Throw %v \nAdvantage: %v \nDisadvantage: %v \nValue: %v", savingThrow.name, savingThrow.hasAdvantage, savingThrow.hasDisadvantage, savingThrow.value)
}
