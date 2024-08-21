package utils

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RetrieveCharacter(characterid string, db *mongo.Collection) (*models.Character, error) {

	objectID, objectIDError := primitive.ObjectIDFromHex(characterid)

	if objectIDError != nil {
		return nil, fmt.Errorf("the error is %s", objectIDError.Error())
	}

	filter := bson.M{"_id": objectID}

	var character models.Character

	charactererror := db.FindOne(context.TODO(), filter).Decode(&character)

	if charactererror != nil {
		return nil, fmt.Errorf("character not found")
	}

	return &character, nil

}

func CreateCharacterToDB(character *models.Character) {
	_, inserterr := database.Characters.InsertOne(context.TODO(), character)
	if inserterr != nil {
		fmt.Println(fmt.Errorf("could not update character: %v", inserterr))
	} else {
		println("Insert of Character was successful")
	}
}

func UpdateCharacterToDB(character *models.Character) {
	filter := bson.D{{Key: "_id", Value: character.ID}}
	update := bson.D{{Key: "$set", Value: character}}
	_, updateerr := database.Characters.UpdateOne(context.TODO(), filter, update)
	if updateerr != nil {
		fmt.Println(fmt.Errorf("could not update character: %v", updateerr))
	} else {
		println("Update of Character was successful")
	}
}
