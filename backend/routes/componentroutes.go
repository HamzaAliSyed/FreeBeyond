package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"
)

func HandleComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/source/create", handlecreatesource)
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

func handlecreatespell(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	utils.OnlyPost(response, request)

}
