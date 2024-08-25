package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func HandleComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/source/create", handlecreatesource)
	mux.HandleFunc("/api/sources/getall", handlegetallsource)
	mux.HandleFunc("/api/spell/create", handlecreatespell)
}

func handlecreatesource(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)

	var Source struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		PublishDate string `json:"publishdate"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&Source)

	if jsonparseerror != nil {
		http.Error(response, "Invalid JSON in request", http.StatusBadRequest)
		return
	}

	var source models.Source

	source.Name = Source.Name
	source.Type = Source.Type
	source.PublishDate = Source.PublishDate

	_, sourceinserterror := database.Sources.InsertOne(context.TODO(), source)

	if sourceinserterror != nil {
		http.Error(response, "Error Inserting Source", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("New Source Created"))
}

func handlegetallsource(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyGet(response, request)

	sourcesCursor, sourcesretrieveerror := database.Sources.Find(context.TODO(), bson.M{})
	if sourcesretrieveerror != nil {
		http.Error(response, "Cannot get sources", http.StatusInternalServerError)
		return
	}

	defer sourcesCursor.Close(context.TODO())

	var sources []string

	for sourcesCursor.Next(context.TODO()) {
		var source models.Source
		sourcedecodeerror := sourcesCursor.Decode(&source)

		if sourcedecodeerror != nil {
			http.Error(response, "Decoding failed", http.StatusInternalServerError)
			return
		}

		sources = append(sources, source.Name)

	}

	if err := sourcesCursor.Err(); err != nil {
		http.Error(response, "Error iterating over sources", http.StatusInternalServerError)
		return
	}

	jsonData, jsonError := json.Marshal(sources)
	if jsonError != nil {
		http.Error(response, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonData)

}

func handlecreatespell(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)

	var SpellCreateRequest struct {
		Name string `json:"name"`
	}

	jsonparseerror := json.NewDecoder(request.Body).Decode(&SpellCreateRequest)

	if jsonparseerror != nil {
		http.Error(response, "Unable To parse JSON", http.StatusBadRequest)
		return
	}

}
