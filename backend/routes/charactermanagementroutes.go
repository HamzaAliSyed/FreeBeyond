package routes

import (
	"backend/models"
	"backend/utils"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleCharacterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/character/create", handleCreateACharacter)
}

func handleCreateACharacter(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyPost(response, request)
	if methodError != nil {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var CharacterRequest struct {
		Name              string            `json:"name"`
		Allignment        string            `json:"allignment"`
		PersonalityTrait  string            `json:"personalitytrait"`
		Ideals            string            `json:"ideals"`
		Bonds             string            `json:"bonds"`
		Flaws             string            `json:"flaws"`
		Strength          int               `json:"strength"`
		Dexterity         int               `json:"dexterity"`
		Constitution      int               `json:"constitution"`
		Intelligence      int               `json:"intelligence"`
		Wisdom            int               `json:"wisdom"`
		Charisma          int               `json:"charisma"`
		PhysicalFeatures  map[string]string `json:"physicalfeatures"`
		CharacterPortrait string            `json:"characterportrait"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&CharacterRequest)

	if jsonParseError != nil {
		http.Error(response, "Error in parsing JSON", http.StatusInternalServerError)
		return
	}

	strippedString := utils.StripBase64Prefix(CharacterRequest.CharacterPortrait)

	characterPortraitData, encodingError := base64.StdEncoding.DecodeString(strippedString)
	if encodingError != nil {
		log.Printf("Base64 decode error: %v", encodingError)
		log.Printf("First 100 chars of input: %s", CharacterRequest.CharacterPortrait[:100])
		http.Error(response, "Invalid image data", http.StatusBadRequest)
		return
	}

	abilityScores := []struct {
		name  string
		value int
	}{
		{"Strength", CharacterRequest.Strength},
		{"Dexterity", CharacterRequest.Dexterity},
		{"Constitution", CharacterRequest.Constitution},
		{"Intelligence", CharacterRequest.Intelligence},
		{"Wisdom", CharacterRequest.Wisdom},
		{"Charisma", CharacterRequest.Charisma},
	}

	var abilityScoreArray []models.AbilityScore

	for _, ability := range abilityScores {
		modifier := (ability.value - 10) / 2
		abilityScoreArray = append(abilityScoreArray, models.AbilityScore{
			AbilityName:     ability.name,
			AbilityScore:    ability.value,
			AbilityModifier: modifier,
		})
	}

	character := models.Character{
		Name:             CharacterRequest.Name,
		CharacterImage:   primitive.Binary{Data: characterPortraitData},
		Allignment:       CharacterRequest.Allignment,
		PersonalityTrait: CharacterRequest.PersonalityTrait,
		Ideals:           CharacterRequest.Ideals,
		Bonds:            CharacterRequest.Bonds,
		Flaws:            CharacterRequest.Flaws,
		ProficiencyBonus: 2,
		AbilityScores:    abilityScoreArray,
		PhysicalFeatures: CharacterRequest.PhysicalFeatures,
	}

	savingThrows := utils.GenerateSavingThrows(abilityScoreArray, character.ProficiencyBonus)
	character.SavingThrows = savingThrows

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&character)
}
