package main

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleCharacterCreation(response http.ResponseWriter, request *http.Request) {

	AllowCorsHeaderAndPreflight(response, request)

	if request.Method != http.MethodPost {
		http.Error(response, "Invalid request method", http.StatusMethodNotAllowed)
		return
	} else {

		var data struct {
			Username string `json:"username"`
		}

		UsernameRetrieveError := json.NewDecoder(request.Body).Decode(&data)

		if UsernameRetrieveError != nil {
			http.Error(response, "Error in parsing data", http.StatusBadRequest)
			return
		}

		var Player struct {
			ID        primitive.ObjectID `bson:"_id"`
			FirstName string             `bson:"first_name"`
			LastName  string             `bson:"last_name"`
		}

		PlayerFullNameRetrivalError := users.FindOne(context.TODO(), bson.M{"username": data.Username}).Decode(&Player)

		if PlayerFullNameRetrivalError != nil {
			http.Error(response, "User not found", http.StatusNotFound)
			return
		}

		//Debugging

		PlayerName := fmt.Sprintf("%s %s", Player.FirstName, Player.LastName)
		fmt.Println(PlayerName)

		NewCharacter := models.Character{
			ID:               primitive.NewObjectID(),
			UserID:           Player.ID,
			PlayerName:       PlayerName,
			ProficiencyBonus: 2,
		}

		CharacterCreated, CreationError := Characters.InsertOne(context.TODO(), NewCharacter)

		if CreationError != nil {
			http.Error(response, "Error creating character", http.StatusInternalServerError)
			return
		}

		newCharacterID := CharacterCreated.InsertedID.(primitive.ObjectID)

		response.WriteHeader(http.StatusCreated)
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(map[string]string{
			"character_id": newCharacterID.Hex(),
		})

	}
}
