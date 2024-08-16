package routes

import (
	"backend/utils"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LanguageHandler struct {
	Langauges *mongo.Collection
}

func (handler *LanguageHandler) HandleCreateLanguages(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methoderror := utils.OnlyPost(response, request); methoderror != nil {
		return
	}

	var LanguagesInstance struct {
		Name            string             `json:"name"`
		Lore            string             `json:"lore"`
		Type            string             `json:"type"`
		TypicalSpeakers primitive.ObjectID `json:"typicalspeakers"`
		Script          string             `json:"script"`
		Source          primitive.ObjectID `json:"source"`
	}

	JsonParseError := json.NewDecoder(request.Body).Decode(&LanguagesInstance)

	if JsonParseError != nil {
		http.Error(response, "Unable to parse json", http.StatusBadRequest)
		return
	}
}
