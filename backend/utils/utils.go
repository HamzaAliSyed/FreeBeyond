package utils

import (
	"math/rand"
	"time"
)

func DieRoller(numberOfDie, sizeOfDie int) []int {
	rand.Seed(time.Now().UnixNano())
	rollResult := make([]int, numberOfDie+1)
	rollTotal := 0
	for index := 1; index <= numberOfDie; index++ {
		roll := rand.Intn(sizeOfDie) + 1
		rollTotal += roll
		rollResult[index] = roll
	}
	rollResult[0] = rollTotal
	return rollResult
}
