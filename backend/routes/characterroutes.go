package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CharacterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/character/create", handleCreateCharacter)
	mux.HandleFunc("/api/character/", handleGetCharacter)
}

func handleCreateCharacter(response http.ResponseWriter, request *http.Request) {

	utils.AllowCorsHeaderAndPreflight(response, request)
	if methodError := utils.OnlyPost(response, request); methodError != nil {
		http.Error(response, methodError.Error(), http.StatusMethodNotAllowed)
		return
	}

	var CharacterRequest struct {
		Name         string `json:"name"`
		Strength     int    `json:"strength"`
		Dexterity    int    `json:"dexterity"`
		Constitution int    `json:"constitution"`
		Intelligence int    `json:"intelligence"`
		Wisdom       int    `json:"wisdom"`
		Charisma     int    `json:"charisma"`
	}

	if jsonParseError := json.NewDecoder(request.Body).Decode(&CharacterRequest); jsonParseError != nil {
		http.Error(response, jsonParseError.Error(), http.StatusBadRequest)
		return
	}

	abilityScores := map[string]int{
		"Strength":     CharacterRequest.Strength,
		"Dexterity":    CharacterRequest.Dexterity,
		"Constitution": CharacterRequest.Constitution,
		"Intelligence": CharacterRequest.Intelligence,
		"Wisdom":       CharacterRequest.Wisdom,
		"Charisma":     CharacterRequest.Charisma,
	}

	var skillToAbility = map[string]string{
		"Acrobatics":      "Dexterity",
		"Animal Handling": "Wisdom",
		"Arcana":          "Intelligence",
		"Athletics":       "Strength",
		"Deception":       "Charisma",
		"History":         "Intelligence",
		"Insight":         "Wisdom",
		"Intimidation":    "Charisma",
		"Investigation":   "Intelligence",
		"Medicine":        "Wisdom",
		"Nature":          "Intelligence",
		"Perception":      "Wisdom",
		"Performance":     "Charisma",
		"Persuasion":      "Charisma",
		"Religion":        "Intelligence",
		"Sleight of Hand": "Dexterity",
		"Stealth":         "Dexterity",
		"Survival":        "Wisdom",
	}

	var abilityScoresArray []models.AbilityScore
	var savingThrowsArray []models.SavingThrow

	for abilityName, score := range abilityScores {
		modifier := (score - 10) / 2

		abilityScore := models.AbilityScore{
			AbilityName: abilityName,
			Score:       score,
			Modifier:    modifier,
		}
		abilityScoresArray = append(abilityScoresArray, abilityScore)

		savingThrow := models.SavingThrow{
			AttributeName:      abilityName,
			Modifer:            modifier,
			AdditionalBonus:    0,
			NumberOProficiency: 0,
			Value:              modifier,
			HasAdvantage:       false,
			HasDisadvantage:    false,
		}
		savingThrowsArray = append(savingThrowsArray, savingThrow)
	}

	var skillsArray []models.Skill

	for skillName, abilityName := range skillToAbility {
		var modifier int
		for _, ability := range abilityScoresArray {
			if ability.AbilityName == abilityName {
				modifier = ability.Modifier
				break
			}
		}

		skill := models.Skill{
			AttributeName:      skillName,
			Modifer:            modifier,
			AdditionalBonus:    0,
			NumberOProficiency: 0,
			Value:              modifier,
			HasAdvantage:       false,
			HasDisadvantage:    false,
		}
		skillsArray = append(skillsArray, skill)
	}

	character := models.Character{
		Name:             CharacterRequest.Name,
		Class:            models.ClassData{}, // Empty for now
		ProficiencyBonus: 2,
		AbilityScores:    abilityScoresArray,
		SavingThrows:     savingThrowsArray,
		Skills:           skillsArray,
	}

	result, err := database.Characters.InsertOne(context.TODO(), character)
	if err != nil {
		http.Error(response, "Error inserting character into database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "Character created successfully",
		"id":      result.InsertedID,
	})
}

func handleGetCharacter(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)

	if methodError := utils.OnlyGet(response, request); methodError != nil {
		http.Error(response, methodError.Error(), http.StatusMethodNotAllowed)
		return
	}

	name := strings.TrimPrefix(request.URL.Path, "/api/character/")
	if name == "" {
		http.Error(response, "Character name is required", http.StatusBadRequest)
		return
	}

	var character *models.Character
	databaseFetchError := database.Characters.FindOne(context.TODO(), bson.M{"name": name}).Decode(&character)
	if databaseFetchError != nil {
		if databaseFetchError == mongo.ErrNoDocuments {
			http.Error(response, "Character doesn't exist", http.StatusNotFound)
		} else {
			http.Error(response, "Database error", http.StatusInternalServerError)
		}
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(character)
}
