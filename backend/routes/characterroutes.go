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
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/create/character", handleCreateCharacter)
	mux.HandleFunc("/api/character/addweaponproficiency", handleAddWeaponProficiencyCharacter)
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

	unArmedAttack := models.NewACBeatingAttack("Unarmed Strike", "Strength", 0, 5, &character, map[models.Damage]string{"Bludgeoning": "1d4"})
	character.AddAttack(unArmedAttack)

	fmt.Printf("Character: %+v\n", character)

	characterInsert := bson.M{
		"name":          character.GetCharacterName(),
		"abilityscores": character.GetAllAbilityScore(),
		"savingthrows":  character.GetAllSavingThrow(),
		"skills":        character.GetAllSkills(),
		"passives":      character.GetAllPassives(),
		"attacks":       character.GetAllAttacks(),
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

func handleAddWeaponProficiencyCharacter(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methodError := utils.OnlyPut(response, request); methodError != nil {
		http.Error(response, methodError.Error(), http.StatusMethodNotAllowed)
		return
	}

	var ProficiencyAddRequestStruct struct {
		Name              string `json:"name"`
		WeaponProficiency string `json:"weaponproficiency"`
	}

	if jsonParseError := json.NewDecoder(request.Body).Decode(&ProficiencyAddRequestStruct); jsonParseError != nil {
		http.Error(response, jsonParseError.Error(), http.StatusBadRequest)
		return
	}

	var character *models.Character
	var characterRetrieveError error
	character, characterRetrieveError = GrabCharacterFromName(ProficiencyAddRequestStruct.Name)
	if characterRetrieveError != nil {
		http.Error(response, characterRetrieveError.Error(), http.StatusInternalServerError)
		return
	}

	character.AddWeaponProficiencies(ProficiencyAddRequestStruct.WeaponProficiency)

	characterUpdate := bson.M{
		"$set": bson.M{
			"weaponproficiencies": character.GetAllWeaponProficiencies(),
		},
	}
	characterFilter := bson.M{
		"_id": character.Id,
	}

	_, updateError := database.Characters.UpdateOne(context.TODO(), characterFilter, characterUpdate)

	if updateError != nil {
		http.Error(response, updateError.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Header().Set("Content-Type", "application/json")
	response.Write([]byte("Weapon Proficiency Added"))

}

func GrabCharacterFromName(characterName string) (*models.Character, error) {
	queryFilter := bson.M{
		"name": characterName,
	}

	var character models.Character

	querySearchError := database.Characters.FindOne(context.TODO(), queryFilter).Decode(&character)

	if querySearchError != nil {
		if querySearchError == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("character not found: %s", characterName)
		}
		if querySearchError == mongo.ErrClientDisconnected {
			return nil, fmt.Errorf("database connection error: %v", querySearchError)
		}
		return nil, fmt.Errorf("error retrieving character: %v", querySearchError)
	}

	return &character, nil
}
