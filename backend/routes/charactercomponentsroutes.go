package routes

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

func HandleComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/components/getabilitymodifier", getAbilityModifier)
}

func getAbilityModifier(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyPost(response, request)
	if methodError != nil {
		http.Error(response, methodError.Error(), http.StatusBadRequest)
		return
	}

	var abilitystruct struct {
		Abilityscore int `json:"abilityscore"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&abilitystruct)

	if jsonParseError != nil {
		http.Error(response, "Unable to Parse JSON", http.StatusBadRequest)
		return
	}

	abilityscoremodifier := (abilitystruct.Abilityscore - 10) / 2

	var responseStruct struct {
		AbilityScoreModifier int `json:"abilityscoremodifier"`
	}

	responseStruct.AbilityScoreModifier = abilityscoremodifier

	jsonResponse, jsonResponseError := json.Marshal(responseStruct)

	if jsonResponseError != nil {
		http.Error(response, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	_, responseWriteError := response.Write(jsonResponse)

	if responseWriteError != nil {
		log.Println("Error writing response:", responseWriteError)
	}

}
