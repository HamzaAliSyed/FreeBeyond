package routes

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
)

func Handleroutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/character/create", handlecharactercreate)
}

func handlecharactercreate(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)
	var CharacterRequest struct {
		Name         string `json:"name"`
		Strength     int    `json:"strength"`
		Dexterity    int    `json:"dexterity"`
		Constitution int    `json:"constitution"`
		Intelligence int    `json:"intelligence"`
		Wisdom       int    `json:"wisdom"`
		Charisma     int    `json:"charisma"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&CharacterRequest)
	if jsonparseerror != nil {
		http.Error(response, "Unable to parse request json", http.StatusBadRequest)
		return
	}

	var character models.Character

	character.SetName(CharacterRequest.Name)
	character.SetStrengthAbilityScore(CharacterRequest.Strength)
	character.SetDexterityAbilityScore(CharacterRequest.Dexterity)
	character.SetConstitutionAbilityScore(CharacterRequest.Constitution)
	character.SetIntelligenceAbilityScore(CharacterRequest.Intelligence)
	character.SetWisdomAbilityScore(CharacterRequest.Wisdom)
	character.SetCharismaAbilityScore(CharacterRequest.Charisma)
	character.SetArmorClass()
	character.SetInitiative(0)
	character.InitializeSavingThrows()
	character.InitializeCharacterSkill()
	utils.CreateCharacterToDB(&character)

}
