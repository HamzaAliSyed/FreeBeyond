package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
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
			FirstName string `bson:"first_name"`
			LastName  string `bson:"last_name"`
		}

		PlayerFullNameRetrivalError := users.FindOne(context.TODO(), bson.M{"username": data.Username}).Decode(&Player)

		if PlayerFullNameRetrivalError != nil {
			http.Error(response, "User not found", http.StatusNotFound)
			return
		}

		//Debugging

		PlayerName := fmt.Sprintf("%s %s", Player.FirstName, Player.LastName)
		fmt.Println(PlayerName)
	}
}
