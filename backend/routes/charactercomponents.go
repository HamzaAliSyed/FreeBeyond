package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CharacterComponentsRoute(mux *http.ServeMux) {
	mux.HandleFunc("/api/charactercomponent/race/create", HandleRaceCreation)
	mux.HandleFunc("/api/charactercomponent/source/create", HandleAddSource)
}

func HandleRaceCreation(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)

	var RaceRequestInstance struct {
		Name               string               `json:"name"`
		MovementSpeed      map[string]int       `json:"movementspeed"`
		Rarity             string               `json:"rarity"`
		Family             string               `json:"family"`
		Size               string               `json:"size"`
		StatsBoost         map[string]int       `json:"statsboost"`
		Languages          []string             `json:"languages"`
		FlavourText        []models.FlavourText `json:"flavourtext"`
		SkillProficiencies []string             `json:"skillproficiencies,omitempty"`
		Attacks            []string             `json:"attacks,omitempty"`
		Spells             []string             `json:"spells,omitempty"`
		Immunities         []string             `json:"immunities,omitempty"`
		Resistances        []string             `json:"resistances,omitempty"`
		PhysicalFeatures   map[string]string    `json:"physicalfeatures"`
		SavingThrows       map[string]string    `json:"savingthrows,omitempty"`
		Source             string               `json:"source"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&RaceRequestInstance)
	if jsonparseerror != nil {
		http.Error(response, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	var SourceLookUp models.Source

	SourceQueryError := database.Sources.FindOne(context.TODO(), bson.M{"name": RaceRequestInstance.Source}).Decode(&SourceLookUp)

	if SourceQueryError != nil {
		if SourceQueryError == mongo.ErrNoDocuments {
			newSource := models.Source{
				Name:        RaceRequestInstance.Source,
				Type:        "",
				PublishDate: "",
			}

			insertResult, insertErr := database.Sources.InsertOne(context.TODO(), newSource)
			if insertErr != nil {
				http.Error(response, "Failed to create new source", http.StatusInternalServerError)
				return
			}

			SourceLookUp.ID = insertResult.InsertedID.(primitive.ObjectID)
			SourceLookUp.Name = newSource.Name

		} else {
			http.Error(response, "Error querying source", http.StatusInternalServerError)
			return
		}
	}

	RaceUpdateDocument := bson.D{
		{Key: "name", Value: RaceRequestInstance.Name},
		{Key: "movementspeed", Value: RaceRequestInstance.MovementSpeed},
		{Key: "rarity", Value: RaceRequestInstance.Rarity},
		{Key: "family", Value: RaceRequestInstance.Family},
		{Key: "size", Value: RaceRequestInstance.Size},
		{Key: "StatsBoost", Value: RaceRequestInstance.StatsBoost},
		{Key: "languages", Value: RaceRequestInstance.Languages},
		{Key: "flavourtext", Value: RaceRequestInstance.FlavourText},
		{Key: "skillproficiencies", Value: RaceRequestInstance.SkillProficiencies},
		{Key: "attacks", Value: RaceRequestInstance.Attacks},
		{Key: "spells", Value: RaceRequestInstance.Spells},
		{Key: "immunities", Value: RaceRequestInstance.Immunities},
		{Key: "resistances", Value: RaceRequestInstance.Resistances},
		{Key: "physicalfeatures", Value: RaceRequestInstance.PhysicalFeatures},
		{Key: "savingthrows", Value: RaceRequestInstance.SavingThrows},
		{Key: "source", Value: SourceLookUp.ID},
	}

	_, RaceInsertError := database.Races.InsertOne(context.TODO(), RaceUpdateDocument)

	if RaceInsertError != nil {
		http.Error(response, "Error creating new race", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("New Race was created and inserted"))
}

func HandleAddSource(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)
	var SourceCreateRequest struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		PublishDate string `json:"publishdate"`
	}

	sourcejsonparseerror := json.NewDecoder(request.Body).Decode(&SourceCreateRequest)
	if sourcejsonparseerror != nil {
		http.Error(response, "Unable to parse json", http.StatusBadRequest)
		return
	}

	var NewSource = bson.D{
		{Key: "name", Value: SourceCreateRequest.Name},
		{Key: "type", Value: SourceCreateRequest.Type},
		{Key: "publishdate", Value: SourceCreateRequest.PublishDate},
	}

	update, updateError := database.Sources.InsertOne(context.TODO(), NewSource)

	if updateError != nil {
		http.Error(response, "Could not update", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(response)
	err := encoder.Encode(update)
	if err != nil {
		http.Error(response, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
