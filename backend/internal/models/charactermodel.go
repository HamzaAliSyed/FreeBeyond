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

func NewCharacter(parameters CharacterParameters) (*Character, error) {
	validName, validNameError := utils.ValidateName(parameters.Name)
	if validNameError != nil {
		return nil, validNameError
	}

	return &Character{
		Name: validName,
	}, nil
}
