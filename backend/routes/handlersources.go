package routes

import (
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SourceHandler struct {
	Sources *mongo.Collection
}

func (handler *SourceHandler) HandleCreateSource(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methoderror := utils.OnlyPost(response, request); methoderror != nil {
		return
	}

	var SourceInstance struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		PublishDate string `json:"publishdate"`
	}

	SourceInstanceParseError := json.NewDecoder(request.Body).Decode(&SourceInstance)

	if SourceInstanceParseError != nil {
		http.Error(response, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	SourceDocument := bson.D{
		{Key: "name", Value: SourceInstance.Name},
		{Key: "type", Value: SourceInstance.Type},
		{Key: "publish date", Value: SourceInstance.PublishDate},
	}

	_, inserterror := handler.Sources.InsertOne(context.TODO(), SourceDocument)

	if inserterror != nil {
		http.Error(response, "Error inserting document", http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("Source created successfully"))

}
