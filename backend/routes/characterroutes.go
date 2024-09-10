package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/create/character", handleCreateCharacter)
}

func handleCreateCharacter(response http.ResponseWriter, request *http.Request) {

	allSkills := map[string]string{
		"Athletics":      "Strength",
		"Acrobatics":     "Dexterity",
		"SleightOfHand":  "Dexterity",
		"Stealth":        "Dexterity",
		"Arcana":         "Intelligence",
		"History":        "Intelligence",
		"Investigation":  "Intelligence",
		"Nature":         "Intelligence",
		"Religion":       "Intelligence",
		"AnimalHandling": "Wisdom",
		"Insight":        "Wisdom",
		"Medicine":       "Wisdom",
		"Perception":     "Wisdom",
		"Survival":       "Wisdom",
		"Deception":      "Charisma",
		"Intimidation":   "Charisma",
		"Performance":    "Charisma",
		"Persuasion":     "Charisma",
	}

	allPassives := map[string]string{
		"Perception":    "Wisdom",
		"Investigation": "Intelligence",
		"Insight":       "Wisdom",
	}

	utils.AllowCorsHeaderAndPreflight(response, request)
	if methodError := utils.OnlyPost(response, request); methodError != nil {
		http.Error(response, methodError.Error(), http.StatusMethodNotAllowed)
		return
	}

	var CreationRequest struct {
		Name         string `json:"name"`
		Strength     int    `json:"strength"`
		Dexterity    int    `json:"dexterity"`
		Constitution int    `json:"constitution"`
		Intelligence int    `json:"intelligence"`
		Wisdom       int    `json:"wisdom"`
		Charisma     int    `json:"charisma"`
	}

	if jsonParseError := json.NewDecoder(request.Body).Decode(&CreationRequest); jsonParseError != nil {
		http.Error(response, jsonParseError.Error(), http.StatusInternalServerError)
		return
	}

	character := models.Character{}
	character = character.SetName(CreationRequest.Name)

	reflectionHandler := reflect.ValueOf(CreationRequest)

	for i := 1; i < 7; i++ {
		stats := reflectionHandler.Field(i)
		statname := reflectionHandler.Type().Field(i).Name
		statvalue := stats.Int()
		abilityScore, abilityScoreCreationError := models.CreateAbilityScore(statname, int(statvalue))
		if abilityScoreCreationError != nil {
			http.Error(response, abilityScoreCreationError.Error(), http.StatusBadRequest)
			return
		}
		character.AddAbilityScoreToCharacter(*abilityScore)
		modifier := (statvalue - 10) / 2
		savingThrow := models.CreateSavingThrow(statname, int(modifier))
		character.AddSavingThrowToCharacter(*savingThrow)
	}

	for skill, stat := range allSkills {
		modifier := character.GetAbilityScoreModifier(stat)
		Skill := models.CreateSkill(skill, stat, modifier)
		character.AddSkillToCharacter(*Skill)
	}

	for skill, stat := range allPassives {
		modifier := character.GetAbilityScoreModifier(stat)
		Passive := models.CreatePassive(skill, stat, modifier)
		character.AddPassiveToCharacter(*Passive)
	}

	fmt.Printf("Character: %+v\n", character)

	characterInsert := bson.M{
		"name":          character.GetCharacterName(),
		"abilityscores": character.GetAllAbilityScore(),
		"savingthrows":  character.GetAllSavingThrow(),
		"skills":        character.GetAllSkills(),
	}

	_, insertError := database.Characters.InsertOne(context.TODO(), characterInsert)

	if insertError != nil {
		http.Error(response, insertError.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte("New Character created"))

}
