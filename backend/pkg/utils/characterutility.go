package utils

import (
	"fmt"
	"strings"
	"unicode"
)

const MaxCharacterNameLength = 30

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
