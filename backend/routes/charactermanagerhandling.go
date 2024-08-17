package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CharacterManagerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/accounts/createacharacter", HandleCharacterCreation)
	mux.HandleFunc("/api/accounts/character/addcharactername", AddCharacterName)
	mux.HandleFunc("/api/accounts/character/addattributes", AddAttributes)
	mux.HandleFunc("/api/charactergeneration/skills/", HandleSkillsFactory)
	mux.HandleFunc("/api/charactergeneration/addcharactermotives", HandleAddCharacterMotives)
}

type AbilityScoreModifiers struct {
	StrengthModifier     int
	DexterityModifier    int
	ConstitutionModifier int
	IntelligenceModifier int
	WisdomModifier       int
	CharismaModifier     int
}

func HandleCharacterCreation(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		return
	}
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

	PlayerFullNameRetrivalError := database.Users.FindOne(context.TODO(), bson.M{"username": data.Username}).Decode(&Player)

	if PlayerFullNameRetrivalError != nil {
		http.Error(response, "User not found", http.StatusNotFound)
		return
	}

	PlayerName := fmt.Sprintf("%s %s", Player.FirstName, Player.LastName)

	NewCharacter := models.Character{
		ID:               primitive.NewObjectID(),
		UserID:           Player.ID,
		PlayerName:       PlayerName,
		ProficiencyBonus: 2,
	}

	CharacterCreated, CreationError := database.Characters.InsertOne(context.TODO(), NewCharacter)

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

func AddCharacterName(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if request.Method != http.MethodPost {
		http.Error(response, "Only Post methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	var CharacterName struct {
		CharacterID   string `json:"characterid"`
		CharacterName string `json:"charactername"`
	}

	CharacterNameRetrieveError := json.NewDecoder(request.Body).Decode(&CharacterName)

	if CharacterNameRetrieveError != nil {
		http.Error(response, "Error finding Character", http.StatusBadRequest)
		return
	}

	CharacterID, CharacterIDConversionError := primitive.ObjectIDFromHex(CharacterName.CharacterID)
	if CharacterIDConversionError != nil {
		http.Error(response, "Error is extracting character id", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": CharacterID}
	update := bson.M{"$set": bson.M{"name": CharacterName.CharacterName}}

	var updatedCharacter models.Character
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := database.Characters.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedCharacter)
	if err != nil {
		http.Error(response, "Error updating character", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]interface{}{
		"status":    "Character name updated successfully",
		"character": updatedCharacter,
	})
}

func AddAttributes(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if MethodError := utils.OnlyPost(response, request); MethodError != nil {
		return
	}

	var CharacterAbilityScore struct {
		CharacterID  string `json:"characterid"`
		Strength     int    `json:"strength"`
		Dexterity    int    `json:"dexterity"`
		Constitution int    `json:"constitution"`
		Intelligence int    `json:"intelligence"`
		Wisdom       int    `json:"wisdom"`
		Charisma     int    `json:"charisma"`
	}

	CharacterAbilityScoreRetrieveError := json.NewDecoder(request.Body).Decode(&CharacterAbilityScore)

	if CharacterAbilityScoreRetrieveError != nil {
		http.Error(response, "Unable to parse request", http.StatusBadRequest)
		return
	}

	targetid := CharacterAbilityScore.CharacterID

	character, characterretrieveerror := utils.RetrieveCharacter(targetid, database.Characters)

	if characterretrieveerror != nil {
		http.Error(response, "Unable to retrieve character", http.StatusBadRequest)
		return
	}

	if character == nil {
		http.Error(response, "Character not found", http.StatusNotFound)
		return
	}

	mainAttributes := models.MainAttributes{
		StrengthScore:     CharacterAbilityScore.Strength,
		DexterityScore:    CharacterAbilityScore.Dexterity,
		ConstitutionScore: CharacterAbilityScore.Constitution,
		IntelligenceScore: CharacterAbilityScore.Intelligence,
		WisdomScore:       CharacterAbilityScore.Wisdom,
		CharismaScore:     CharacterAbilityScore.Charisma,
	}

	character.MainAttributes = mainAttributes

	character.Modifiers = models.Modifiers{
		StrengthModifier:     utils.ModifierCalculator(CharacterAbilityScore.Strength),
		DexterityModifier:    utils.ModifierCalculator(CharacterAbilityScore.Dexterity),
		ConstitutionModifier: utils.ModifierCalculator(CharacterAbilityScore.Constitution),
		IntelligenceModifier: utils.ModifierCalculator(CharacterAbilityScore.Intelligence),
		WisdomModifier:       utils.ModifierCalculator(CharacterAbilityScore.Wisdom),
		CharismaModifier:     utils.ModifierCalculator(CharacterAbilityScore.Charisma),
	}

	utils.InitialSavingThrowsGenerator(character)
	utils.InitializeSkillsArray(character, database.Skills)

	filter := bson.M{"_id": character.ID}
	update := bson.M{
		"$set": bson.M{
			"mainattributes": character.MainAttributes,
			"modifiers":      character.Modifiers,
			"savingthrow":    character.SavingThrow,
			"skills":         character.Skills,
		},
	}

	_, updateErr := database.Characters.UpdateOne(context.TODO(), filter, update)
	if updateErr != nil {
		http.Error(response, "Failed to update character", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{
		"status": "Character attributes updated successfully",
	})
}
func HandleSkillsFactory(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Skill Factory called")
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methoderror := utils.OnlyPost(response, request); methoderror != nil {
		return
	}

	var SkillInstance struct {
		Name                string `json:"name"`
		AssociatedAttribute string `json:"associatedattribute"`
	}

	SkillRetrieveError := json.NewDecoder(request.Body).Decode(&SkillInstance)

	if SkillRetrieveError != nil {
		http.Error(response, "Invalid Attribute", http.StatusBadRequest)
		return
	}
	SkillDocument := bson.D{
		{Key: "name", Value: SkillInstance.Name},
		{Key: "associatedattribute", Value: SkillInstance.AssociatedAttribute},
	}

	_, insertErr := database.Skills.InsertOne(context.TODO(), SkillDocument)

	if insertErr != nil {
		http.Error(response, "Failed to insert skill", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("Skill inserted successfully"))

}

func HandleAddCharacterMotives(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methoderror := utils.OnlyPost(response, request); methoderror != nil {
		return
	}

	var CharacterMotiveInstance struct {
		CharacterID       string `json:"characterid"`
		PersonalityTraits string `json:"personalitytraits"`
		Ideals            string `json:"ideals"`
		Bonds             string `json:"bonds"`
		Flaws             string `json:"flaws"`
	}

	CharacterMotiveExtractError := json.NewDecoder(request.Body).Decode(&CharacterMotiveInstance)
	if CharacterMotiveExtractError != nil {
		http.Error(response, "Bad request in JSON", http.StatusBadRequest)
	}

	character, characterRetireError := utils.RetrieveCharacter(CharacterMotiveInstance.CharacterID, database.Characters)

	if characterRetireError != nil {
		http.Error(response, "Cannot find the character", http.StatusBadRequest)
		return
	}

	character.CharacterMotives = models.CharacterMotives{
		PersonalityTraits: CharacterMotiveInstance.PersonalityTraits,
		Bonds:             CharacterMotiveInstance.Bonds,
		Ideals:            CharacterMotiveInstance.Ideals,
		Flaws:             CharacterMotiveInstance.Flaws,
	}

	filter := bson.M{"_id": character.ID}
	update := bson.M{
		"$set": bson.M{
			"charactermotives": character.CharacterMotives,
		},
	}

	_, updateErr := database.Characters.UpdateOne(context.TODO(), filter, update)
	if updateErr != nil {
		http.Error(response, "Failed to update character", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(map[string]string{
		"status": "Character motives updated successfully",
	})

}
