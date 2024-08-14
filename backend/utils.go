package main

import (
	"backend/models"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllowCorsHeaderAndPreflight(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
		return
	}

}

func OnlyPost(response http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodPost {
		http.Error(response, "Only POST method allowed on the end point", http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	return nil
}

func RetrieveCharacter(characterid string) (*models.Character, error) {

	objectID, objectIDError := primitive.ObjectIDFromHex(characterid)

	if objectIDError != nil {
		return nil, fmt.Errorf("invalid ID format")
	}

	filter := bson.M{"_id": objectID}

	var character models.Character

	charactererror := Characters.FindOne(context.TODO(), filter).Decode(&character)

	if charactererror != nil {
		return nil, fmt.Errorf("character not found")
	}

	return &character, nil

}

func AttributeModifierCalculator(statvalue int) int {
	statmodifier := (statvalue - 10) / 2
	return statmodifier
}
