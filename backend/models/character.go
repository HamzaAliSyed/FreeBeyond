package models

import (
	"backend/database"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterOperations interface{}

type AbilityScore struct {
	StatName     string `bson:"statname"`
	StatValue    int    `bson:"statValue"`
	StatModifier int    `bson:"statmodifier"`
}

type Character struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	charactername string             `bson:"charactername,omitempty"`
	abilityscores []AbilityScore     `bson:"abilityscores,omitempty"`
}

func (character Character) SetName(name string) Character {
	character.charactername = name
	return character
}

func (character *Character) GetName(id string) (string, error) {
	queryfilter := bson.M{
		"_id": id,
	}

	querySearchCharacterError := database.Characters.FindOne(context.TODO(), queryfilter).Decode(&character)
	if querySearchCharacterError != nil {
		return "", querySearchCharacterError
	}

	return character.charactername, nil

}

func CreateAbilityScore(statname string, statvalue int) (*AbilityScore, error) {
	var abilityscore AbilityScore
	if statvalue < 3 || statvalue > 18 {
		return nil, errors.New("invalid rolled stat, stat value must be between 3 and 18")
	}

	abilityscore.StatName = statname
	abilityscore.StatValue = statvalue
	abilityscore.StatModifier = (statvalue - 10) / 2
	return &abilityscore, nil
}

func (character *Character) AddAbilityScoreToCharacter(abilityscore AbilityScore) error {
	character.abilityscores = append(character.abilityscores, abilityscore)
	return nil
}

func (character *Character) GetCharacterName() string {
	return character.charactername
}

func (character *Character) GetAllAbilityScore() []AbilityScore {
	return character.abilityscores
}
