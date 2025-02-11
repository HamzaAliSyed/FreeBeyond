package utils

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

const MaxCharacterNameLength = 30
const LowerAbilityScoreValue = 1
const HigherAbilityScoreValue = 30
const NewCharacterHigherAbilityScoreValue = 18

func ValidateName(name string) (string, error) {
	trimmedName := strings.TrimSpace(name)

	if len([]rune(trimmedName)) > MaxCharacterNameLength {
		return "", fmt.Errorf("provided name exceeds the maximum character limit of %d", MaxCharacterNameLength)
	}

	for _, letter := range trimmedName {
		if !unicode.IsLetter(letter) && !unicode.IsSpace(letter) {
			return "", fmt.Errorf("%c is not a valid unicode letter", letter)
		}
	}

	normalizedName := capitalizeNameFirstLetters(trimmedName)
	return normalizedName, nil
}

func capitalizeNameFirstLetters(fullName string) string {
	names := strings.Fields(fullName)
	for index, name := range names {
		names[index] = func() string {
			if name == "" {
				return ""
			}

			letters := []rune(name)
			letters[0] = unicode.ToUpper(letters[0])
			for secondindex := 1; secondindex < len(letters); secondindex++ {
				letters[secondindex] = unicode.ToLower(letters[secondindex])
			}

			return string(letters)
		}()
	}

	return strings.Join(names, " ")
}

func NewCharacterValidateScoreAndGenerateModifier(score int) (int, int, error) {
	if score >= LowerAbilityScoreValue && score <= NewCharacterHigherAbilityScoreValue {
		modifierValue := modifierGenerator(score)
		return score, modifierValue, nil
	} else {
		return 0, 0, fmt.Errorf("score cannot be more than 18 or less than 1")
	}
}

func ValidateScoreAndGenerateModifier(score int) (int, int, error) {
	if score >= LowerAbilityScoreValue && score <= HigherAbilityScoreValue {
		modifierValue := modifierGenerator(score)
		return score, modifierValue, nil
	} else {
		return 0, 0, fmt.Errorf("score cannot be more than 18 or less than 1")
	}
}

func modifierGenerator(score int) int {
	floatValue := float64(score)
	floatValue = (floatValue - 10.0) / 2
	floatValue = math.Floor(floatValue)
	return int(floatValue)
}
