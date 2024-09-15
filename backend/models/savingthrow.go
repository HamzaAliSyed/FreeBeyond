package models

type SavingThrow struct {
	savingThrowName                  string  `bson:"savingthrowname"`
	savingThrowMod                   int     `bson:"savingthrowmod"`
	savingThrowNumberOfProficiencies float64 `bson:"savingthrownumberofproficiencies"`
	additionalBonus                  int     `bson:"additionalbonus"`
	hasAdvantage                     bool    `bson:"hasadvantage"`
	hasDisadvantage                  bool    `bson:"hasdisadvantage"`
	value                            int     `bson:"value"`
}

func (savingThrow *SavingThrow) SetSavingThrowName(name string) {
	savingThrow.savingThrowName = name
}

func (savingThrow *SavingThrow) GetSavingThrowName() string {
	return savingThrow.savingThrowName
}

func (savingThrow *SavingThrow) SetSavingThrowMod(value int) {
	savingThrow.savingThrowMod = value
}

func (savingThrow *SavingThrow) UpdateThrowNumberOfProficiency(value float64) {
	savingThrow.savingThrowNumberOfProficiencies += value
}

func (savingThrow *SavingThrow) SetSavingThrowAdditionalBonus(value int) {
	savingThrow.additionalBonus = value
}

func (savingThrow *SavingThrow) ToggleSavingThrowAdvantage(isTrue bool) {
	savingThrow.hasAdvantage = isTrue
}

func (savingThrow *SavingThrow) ToggleSavingThrowDisAdvantage(isTrue bool) {
	savingThrow.hasDisadvantage = isTrue
}

func (savingThrow *SavingThrow) SetSavingThrowValue(valueFinal int) {
	savingThrow.value = valueFinal
}

func (savingThrow *SavingThrow) GetSavingThrowValue() int {
	return savingThrow.value
}

func (savingThrow *SavingThrow) CreateSavingThrow(name string, mod int) {
	savingThrow.SetSavingThrowName(name)
	savingThrow.SetSavingThrowMod(mod)
	savingThrow.savingThrowNumberOfProficiencies = 0.0
	savingThrow.ToggleSavingThrowAdvantage(false)
	savingThrow.ToggleSavingThrowDisAdvantage(false)
	savingThrow.SetSavingThrowValue(mod)
}

func (savingThrow *SavingThrow) CalculateValue(character *Character) int {
	value := savingThrow.additionalBonus + int((savingThrow.savingThrowNumberOfProficiencies)*float64(character.proficiencyBonus)) + savingThrow.savingThrowMod
	return value
}
