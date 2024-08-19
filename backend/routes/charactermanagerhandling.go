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

const NONLEGENDARYCHARACTERMAX = 20
const LEGENDARYCHARACTERMAX = 26

var InstanceMainAttributes models.MainAttributes
var InstanceModifiers models.Modifiers
var InstanceSavingThrow []models.SavingThrow
var InstanceSkill models.Skills

func CharacterManagerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/accounts/createacharacter", HandleCharacterCreation)
	mux.HandleFunc("/api/accounts/character/addcharactername", AddCharacterName)
	mux.HandleFunc("/api/accounts/character/addattributes", AddAttributes)
	mux.HandleFunc("/api/charactergeneration/skills/", HandleSkillsFactory)
	mux.HandleFunc("/api/charactergeneration/addcharactermotives", HandleAddCharacterMotives)
	mux.HandleFunc("/api/charactermodify/addfeats", HandleAddFeatsToCharacter)
	mux.HandleFunc("/api/charactermodify/addbackground", AddBackGroundCharacter)
	mux.HandleFunc("/api/character/addrace/", HandleAddRacetoCharacter)
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
	} else {
		fmt.Printf("Character had kick ass stats added and i am here to debug %v", character.ID)
		var characteridstring string = character.ID.Hex()
		maxcarryweight := utils.MaxCarryWeightCalculator(characteridstring)
		carryweight := utils.CarryWeightCalculator(characteridstring)
		fmt.Println("Now trying second update")
		character.MaxCarryWeight = maxcarryweight
		character.CarryWeight = carryweight
		_, replaceError := database.Characters.ReplaceOne(context.TODO(), filter, character)
		if replaceError != nil {
			fmt.Println("Something went wrong in the second update")
			msg := fmt.Errorf("error happened: %v", replaceError)
			fmt.Println(msg)
		} else {
			fmt.Println("Carry Weight problem solved")
		}
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

func HandleAddFeatsToCharacter(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		return
	}
	var CharacterFeatInstance struct {
		FeatsID     primitive.ObjectID `json:"featsid"`
		CharacterID primitive.ObjectID `json:"characterid"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&CharacterFeatInstance)
	if jsonparseerror != nil {
		http.Error(response, "Cannot parse json error", http.StatusBadRequest)
		return
	}

	characteridstring := CharacterFeatInstance.CharacterID.Hex()
	featsidstring := CharacterFeatInstance.FeatsID.Hex()

	character, characterretrieveerror := utils.RetrieveCharacter(characteridstring, database.Characters)
	if characterretrieveerror != nil {
		http.Error(response, "Invalid Character ID", http.StatusInternalServerError)
		return
	}

	feats, featserrorretrieve := utils.RetrieveFeats(featsidstring, database.Feats)
	if featserrorretrieve != nil {
		http.Error(response, "Invalid Feat ID", http.StatusInternalServerError)
		return
	}

	switch feats.Prerequisite {
	case "None":
		{
			charactermodificationarray := feats.CharacterModifications
			for _, modification := range charactermodificationarray {
				if modification.Category == "Ability Score" {
					switch modification.Attribute {
					case "Strength":
						{
							if character.MainAttributes.StrengthScore >= NONLEGENDARYCHARACTERMAX {
								http.Error(response, "Character already on max strength", http.StatusForbidden)
								return
							} else {
								character.MainAttributes.StrengthScore += modification.Value
								character.Modifiers.StrengthModifier = utils.ModifierCalculator(character.MainAttributes.StrengthScore)
								InstanceSavingThrow = utils.InitialSavingThrowsGenerator(character)
								InstanceSkill = utils.InitializeSkillsArray(character, database.Skills)
							}
						}
					case "Dexterity":
						{
							if character.MainAttributes.DexterityScore >= NONLEGENDARYCHARACTERMAX {
								http.Error(response, "Character already on max dexterity", http.StatusForbidden)
								return
							} else {
								character.MainAttributes.DexterityScore += modification.Value
								character.Modifiers.DexterityModifier = utils.ModifierCalculator(character.MainAttributes.DexterityScore)
								InstanceSavingThrow = utils.InitialSavingThrowsGenerator(character)
								InstanceSkill = utils.InitializeSkillsArray(character, database.Skills)
							}
						}
					case "Constitution":
						{
							if character.MainAttributes.ConstitutionScore >= NONLEGENDARYCHARACTERMAX {
								http.Error(response, "Character already on max constitution", http.StatusForbidden)
								return
							} else {
								character.MainAttributes.ConstitutionScore += modification.Value
								character.Modifiers.ConstitutionModifier = utils.ModifierCalculator(character.MainAttributes.ConstitutionScore)
								InstanceSavingThrow = utils.InitialSavingThrowsGenerator(character)
								InstanceSkill = utils.InitializeSkillsArray(character, database.Skills)
							}
						}
					case "Intelligence":
						{
							if character.MainAttributes.IntelligenceScore >= NONLEGENDARYCHARACTERMAX {
								http.Error(response, "Character already on max intelligence", http.StatusForbidden)
								return
							} else {
								character.MainAttributes.IntelligenceScore += modification.Value
								character.Modifiers.IntelligenceModifier = utils.ModifierCalculator(character.MainAttributes.IntelligenceScore)
								InstanceSavingThrow = utils.InitialSavingThrowsGenerator(character)
								InstanceSkill = utils.InitializeSkillsArray(character, database.Skills)
							}
						}
					case "Wisdom":
						{
							if character.MainAttributes.WisdomScore >= NONLEGENDARYCHARACTERMAX {
								http.Error(response, "Character already on max wisdom", http.StatusForbidden)
								return
							} else {
								character.MainAttributes.WisdomScore += modification.Value
								character.Modifiers.WisdomModifier = utils.ModifierCalculator(character.MainAttributes.WisdomScore)
								InstanceSavingThrow = utils.InitialSavingThrowsGenerator(character)
								InstanceSkill = utils.InitializeSkillsArray(character, database.Skills)
							}
						}
					case "Charisma":
						{
							if character.MainAttributes.CharismaScore >= NONLEGENDARYCHARACTERMAX {
								http.Error(response, "Character already on max charisma", http.StatusForbidden)
								return
							} else {
								character.MainAttributes.CharismaScore += modification.Value
								character.Modifiers.CharismaModifier = utils.ModifierCalculator(character.MainAttributes.CharismaScore)
								InstanceSavingThrow = utils.InitialSavingThrowsGenerator(character)
								InstanceSkill = utils.InitializeSkillsArray(character, database.Skills)
							}
						}
					}
				} else if modification.Category == "Saving Throws" {
					InstanceSavingThrow = character.SavingThrow
					for i := range InstanceSavingThrow {
						switch InstanceSavingThrow[i].Attribute {
						case "Strength":
							{
								InstanceSavingThrow[i].SavingThrowValue += modification.Value
							}
						case "Dexterity":
							{
								InstanceSavingThrow[i].SavingThrowValue += modification.Value
							}
						case "Constitution":
							{
								InstanceSavingThrow[i].SavingThrowValue += modification.Value
							}
						case "Intelligence":
							{
								InstanceSavingThrow[i].SavingThrowValue += modification.Value
							}
						case "Wisdom":
							{
								InstanceSavingThrow[i].SavingThrowValue += modification.Value
							}
						case "Charisma":
							{
								InstanceSavingThrow[i].SavingThrowValue += modification.Value
							}
						}
					}
					character.SavingThrow = InstanceSavingThrow
				} else if modification.Category == "Skills" {

				}
			}

			character.SavingThrow = InstanceSavingThrow
			character.Skills = InstanceSkill
			character.Feats = append(character.Feats, *feats)

			updateCharacter := bson.M{
				"$set": character,
			}

			_, updateerror := database.Characters.UpdateOne(
				context.TODO(),
				bson.M{"_id": character.ID},
				updateCharacter,
			)

			if updateerror != nil {
				fmt.Printf("Error updating character: %v\n", updateerror)
				http.Error(response, "Failed to add feat to character", http.StatusInternalServerError)
				return
			}

			response.WriteHeader(http.StatusOK)
			response.Write([]byte("Feat was added to the character"))
		}
	default:
		{
			http.Error(response, "Feat cannot be added because it does not meet prerequisites", http.StatusUnauthorized)
			return
		}
	}
}

func AddBackGroundCharacter(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)
	var BackgroundInstance struct {
		CharacterID     string `json:"characterID"`
		BackgroundName  string `json:"backgroundname"`
		FirstAttribute  string `json:"firstattribute"`
		SecondAttribute string `json:"secondattribute"`
		FirstValue      int    `json:"firstvalue"`
		SecondValue     int    `json:"secondvalue"`
		FirstSkill      string `json:"firstskill"`
		SecondSkill     string `json:"secondskill"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&BackgroundInstance)

	if jsonparseerror != nil {
		http.Error(response, "Error Parsing the request JSON", http.StatusBadRequest)
		return
	}

	character, characterErr := utils.RetrieveCharacter(BackgroundInstance.CharacterID, database.Characters)
	if characterErr != nil {
		http.Error(response, "Error finding Character", http.StatusBadRequest)
		return
	}

	characterID := BackgroundInstance.CharacterID

	capFirstAttribute := utils.CapitalizeFirstLetter(BackgroundInstance.FirstAttribute)
	capSecondAttribute := utils.CapitalizeFirstLetter(BackgroundInstance.SecondAttribute)
	fmt.Println("First we are updating first attribute")

	character.MainAttributes, character.Modifiers = utils.UpdateAbilityScore(characterID, BackgroundInstance.FirstAttribute, BackgroundInstance.FirstValue)
	utils.UpdateCharacterToDB(character)

	character.SavingThrow = utils.UpdateSavingThrowsAfterASI(character)
	character.Skills = utils.UpdateSkillsAfterASI(character, capFirstAttribute)
	character.MaxCarryWeight = utils.UpdateMaxCarryWeight(character)
	utils.UpdateCharacterToDB(character)

	character.MainAttributes, character.Modifiers = utils.UpdateAbilityScore(characterID, BackgroundInstance.SecondAttribute, BackgroundInstance.SecondValue)
	utils.UpdateCharacterToDB(character)
	character.Skills = utils.UpdateSkillsAfterASI(character, capSecondAttribute)
	character.SavingThrow = utils.UpdateSavingThrowsAfterASI(character)
	character.MaxCarryWeight = utils.UpdateMaxCarryWeight(character)
	character.Background = BackgroundInstance.BackgroundName
	utils.UpdateCharacterToDB(character)

	utils.AddProfiencyToSkill(character, BackgroundInstance.FirstSkill)
	utils.AddProfiencyToSkill(character, BackgroundInstance.SecondSkill)
	utils.UpdateCharacterToDB(character)

	// We are now returning
	fmt.Println("Now returning response")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Background added to character"))
}

func HandleAddRacetoCharacter(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)

	var RaceAddRequest struct {
		CharacterID    string   `json:"characterid"`
		RaceName       string   `json:"racename"`
		ExtraLanguages []string `json:"language"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&RaceAddRequest)

	if jsonparseerror != nil {
		http.Error(response, "Error Parsing JSON", http.StatusBadRequest)
		return
	}

	character, characterretrieveerror := utils.RetrieveCharacter(RaceAddRequest.CharacterID, database.Characters)

	if characterretrieveerror != nil {
		http.Error(response, characterretrieveerror.Error(), http.StatusBadRequest)
		return
	}

	if character.Race != "" {
		http.Error(response, "The character already has a race", http.StatusUnauthorized)
		return
	}

	var targetRace models.Race

	racename := bson.D{
		{Key: "name", Value: RaceAddRequest.RaceName},
	}

	raceQueryError := database.Races.FindOne(context.TODO(), racename).Decode(&targetRace)

	if raceQueryError != nil {
		http.Error(response, raceQueryError.Error(), http.StatusBadRequest)
		return
	}
	for speedType, speedValue := range targetRace.MovementSpeed {
		switch speedType {
		case "landspeed":
			{
				character.LandSpeed = speedValue
			}
		case "flyingspeed":
			{
				character.FlyingSpeed = speedValue
			}
		case "swimmingspeed":
			{
				character.SwimmingSpeed = speedValue
			}
		case "climbingspeed":
			{
				character.ClimbingSpeed = speedValue
			}
		case "burrowingspeed":
			{
				character.BurrowingSpeed = speedValue
			}
		}
	}

	if character.Appearance == nil {
		character.Appearance = make(map[string]string)
		character.Appearance["size"] = targetRace.Size
	} else {
		character.Appearance["size"] = targetRace.Size
	}

	utils.UpdateCharacterToDB(character)

	characterID := RaceAddRequest.CharacterID

	for stat, boost := range targetRace.StatsBoost {
		character.MainAttributes, character.Modifiers = utils.UpdateAbilityScore(characterID, stat, boost)
		utils.UpdateCharacterToDB(character)
		statWithFirstCap := utils.CapitalizeFirstLetter(stat)
		character.Skills = utils.UpdateSkillsAfterASI(character, statWithFirstCap)
		character.SavingThrow = utils.UpdateSavingThrowsAfterASI(character)
		character.MaxCarryWeight = utils.UpdateMaxCarryWeight(character)
		utils.UpdateCharacterToDB(character)
	}

	languages := append(targetRace.Languages, RaceAddRequest.ExtraLanguages...)
	character.Languages = languages

	utils.UpdateCharacterToDB(character)

	if len(targetRace.AdvantageSkills) > 0 {
		for _, skill := range targetRace.AdvantageSkills {
			utils.AddAdvantageToSkill(character, skill)
		}
	}

	if len(targetRace.SkillProficiencies) > 0 {
		for _, skill := range targetRace.SkillProficiencies {
			utils.AddProfiencyToSkill(character, skill)
		}
	}

	utils.UpdateCharacterToDB(character)

	if len(targetRace.SavingThrows) > 0 {
		for operation, savingthrow := range targetRace.SavingThrows {
			if operation == "Advantage" {
				character.SavingThrow = utils.AddAdvantageToSavingThrows(character, savingthrow)
			} else if operation == "Add Proficiency" {
				character.SavingThrow = utils.AddProfiencyToSavingThrow(character, savingthrow)
			}
		}
	}

	utils.UpdateCharacterToDB(character)

}
