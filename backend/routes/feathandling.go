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

func FeatRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/feats/create", HandleCreateFeat)
}

func HandleCreateFeat(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		return
	}

	var FeatFromResponse struct {
		Name             string                                      `json:"name"`
		Description      string                                      `json:"description"`
		Prerequisite     string                                      `json:"prerequisite"`
		Feature          []models.FlavourText                        `json:"flavourfeatures"`
		AbilityModifider []models.CharacterStatsAndAbilitiesModifier `json:"abilitiesmodifier"`
		Charges          []models.ChargeBasedAbilities               `json:"chargeability"`
		Attack           []models.AnAttack                           `json:"attack"`
		Spells           []models.SpellAttack                        `json:"spells"`
		Source           string                                      `json:"Source"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&FeatFromResponse)

	if jsonparseerror != nil {
		http.Error(response, "Unable to parse the json in request", http.StatusBadRequest)
		return
	}

	var SourceLookUp models.Source

	SourceQueryError := database.Source.FindOne(context.TODO(), bson.M{"name": FeatFromResponse.Source}).Decode(&SourceLookUp)

	if SourceQueryError != nil {
		if SourceQueryError == mongo.ErrNoDocuments {
			newSource := models.Source{
				Name:        FeatFromResponse.Source,
				Type:        "",
				PublishDate: "",
			}

			insertResult, insertErr := database.Source.InsertOne(context.TODO(), newSource)
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

	FeatBSONDocument := bson.D{
		{Key: "featname", Value: FeatFromResponse.Name},
		{Key: "featdescription", Value: FeatFromResponse.Description},
		{Key: "prequisite", Value: FeatFromResponse.Prerequisite},
		{Key: "textfeature", Value: FeatFromResponse.Feature},
		{Key: "charactermodification", Value: FeatFromResponse.AbilityModifider},
		{Key: "chargeabilities", Value: FeatFromResponse.Charges},
		{Key: "attack", Value: FeatFromResponse.Attack},
		{Key: "spells", Value: FeatFromResponse.Spells},
		{Key: "Source", Value: SourceLookUp.ID},
	}

	_, featinserterror := database.Feats.InsertOne(context.TODO(), FeatBSONDocument)

	if featinserterror != nil {
		http.Error(response, "Error inserting the document", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("Successfully created new feat"))

}
