package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/components/getabilitymodifier", getAbilityModifier)
	mux.HandleFunc("/api/components/createsource", handlecreatesource)
	mux.HandleFunc("/api/components/getallsources", getAllSources)
	mux.HandleFunc("/api/components/getallsourcesnames", getAllSourcesNames)
}

func getAbilityModifier(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyPost(response, request)
	if methodError != nil {
		http.Error(response, methodError.Error(), http.StatusBadRequest)
		return
	}

	var abilitystruct struct {
		Abilityscore int `json:"abilityscore"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&abilitystruct)

	if jsonParseError != nil {
		http.Error(response, "Unable to Parse JSON", http.StatusBadRequest)
		return
	}

	abilityscoremodifier := (abilitystruct.Abilityscore - 10) / 2

	var responseStruct struct {
		AbilityScoreModifier int `json:"abilityscoremodifier"`
	}

	responseStruct.AbilityScoreModifier = abilityscoremodifier

	jsonResponse, jsonResponseError := json.Marshal(responseStruct)

	if jsonResponseError != nil {
		http.Error(response, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	_, responseWriteError := response.Write(jsonResponse)

	if responseWriteError != nil {
		log.Println("Error writing response:", responseWriteError)
	}

}

func handlecreatesource(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		http.Error(response, methoderror.Error(), http.StatusBadRequest)
		return
	}

	var SourceStruct struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		PublishDate string `json:"publishdate"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&SourceStruct)

	if jsonParseError != nil {
		http.Error(response, "Unable to Parse JSON", http.StatusBadRequest)
		return
	}

	publishDate, err := time.Parse("January-02-2006", SourceStruct.PublishDate)
	if err != nil {
		http.Error(response, "Invalid date format. Use 'Month-Day-Year' format.", http.StatusBadRequest)
		return
	}

	newSource := models.Source{
		Name:        SourceStruct.Name,
		Type:        SourceStruct.Type,
		PublishDate: publishDate,
	}

	_, sourceinserterror := database.Sources.InsertOne(context.TODO(), newSource)

	if sourceinserterror != nil {
		http.Error(response, "Error Inserting Source", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("New Source Created"))

}

func getAllSources(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyGet(response, request)

	if methodError != nil {
		http.Error(response, "Only Get Method allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, cursorError := database.Sources.Find(context.TODO(), bson.M{})

	if cursorError != nil {
		http.Error(response, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	var sourcesQuery []bson.M

	if cursorQueryAllError := cursor.All(context.TODO(), &sourcesQuery); cursorQueryAllError != nil {
		http.Error(response, "Failed to decode data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sourcesQuery)
}

func getAllSourcesNames(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyGet(response, request)

	if methodError != nil {
		http.Error(response, "Only Get Method allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, cursorError := database.Sources.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"name": 1}))

	if cursorError != nil {
		http.Error(response, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var sourceNames []string

	for cursor.Next(context.TODO()) {
		var result struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&result); err != nil {
			http.Error(response, "Failed to decode data", http.StatusInternalServerError)
			return
		}
		sourceNames = append(sourceNames, result.Name)
	}

	if err := cursor.Err(); err != nil {
		http.Error(response, "Cursor error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sourceNames)
}
