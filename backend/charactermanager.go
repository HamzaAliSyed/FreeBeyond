package main

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AbilityScoreModifiers struct {
	StrengthModifier     int
	DexterityModifier    int
	ConstitutionModifier int
	IntelligenceModifier int
	WisdomModifier       int
	CharismaModifier     int
}

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
			ID        primitive.ObjectID `bson:"_id"`
			FirstName string             `bson:"first_name"`
			LastName  string             `bson:"last_name"`
		}

		PlayerFullNameRetrivalError := users.FindOne(context.TODO(), bson.M{"username": data.Username}).Decode(&Player)

		if PlayerFullNameRetrivalError != nil {
			http.Error(response, "User not found", http.StatusNotFound)
			return
		}

		//Debugging

		PlayerName := fmt.Sprintf("%s %s", Player.FirstName, Player.LastName)
		fmt.Println(PlayerName)

		NewCharacter := models.Character{
			ID:               primitive.NewObjectID(),
			UserID:           Player.ID,
			PlayerName:       PlayerName,
			ProficiencyBonus: 2,
		}

		CharacterCreated, CreationError := Characters.InsertOne(context.TODO(), NewCharacter)

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
}

func AddCharacterName(response http.ResponseWriter, request *http.Request) {
	AllowCorsHeaderAndPreflight(response, request)
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

	err := Characters.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedCharacter)
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
	AllowCorsHeaderAndPreflight(response, request)
	if MethodError := OnlyPost(response, request); MethodError != nil {
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

	character, characterretrieveerror := RetrieveCharacter(targetid)

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
		StrengthModifier:     ModifierCalculator(CharacterAbilityScore.Strength),
		DexterityModifier:    ModifierCalculator(CharacterAbilityScore.Dexterity),
		ConstitutionModifier: ModifierCalculator(CharacterAbilityScore.Constitution),
		IntelligenceModifier: ModifierCalculator(CharacterAbilityScore.Intelligence),
		WisdomModifier:       ModifierCalculator(CharacterAbilityScore.Wisdom),
		CharismaModifier:     ModifierCalculator(CharacterAbilityScore.Charisma),
	}

	InitialSavingThrowsGenerator(character)

	filter := bson.M{"_id": character.ID}
	update := bson.M{
		"$set": bson.M{
			"mainattributes": character.MainAttributes,
			"modifiers":      character.Modifiers,
			"savingthrow":    character.SavingThrow,
		},
	}

	_, updateErr := Characters.UpdateOne(context.TODO(), filter, update)
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
	AllowCorsHeaderAndPreflight(response, request)
	OnlyPost(response, request)

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

	_, insertErr := skills.InsertOne(context.TODO(), SkillDocument)

	if insertErr != nil {
		http.Error(response, "Failed to insert skill", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("Skill inserted successfully"))

}
