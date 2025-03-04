package models

import (
	"backend/pkg/utils"
)

type Character struct {
	Name string
}

type CharacterParameters struct {
	Name string
}

func NewCharacter(parameters CharacterParameters) (Character, error) {
	validName, validNameError := utils.ValidateName(parameters.Name)
	var character Character
	if validNameError != nil {
		return character, validNameError
	}

	character.Name = validName

	return character, nil
}
