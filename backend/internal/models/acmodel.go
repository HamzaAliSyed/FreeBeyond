package models

import "fmt"

type AC struct {
	finalValue int
}

func (ac *AC) CalculateAC(dexMod int) {
	ac.finalValue = 10 + dexMod
}

func (ac AC) String() string {
	return fmt.Sprintf("AC: %d", ac.finalValue)
}
