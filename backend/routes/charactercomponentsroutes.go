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
	mux.HandleFunc("/api/components/addspells", handleAddSpells)
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

func handleAddSpells(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyPost(response, request)

	if methodError != nil {
		http.Error(response, "Only Post Method allowed", http.StatusMethodNotAllowed)
		return
	}

	var SpellRequest struct {
		Name          string            `json:"name"`
		Level         int               `json:"level"`
		CastingTime   string            `json:"castingtime"`
		Duration      string            `json:"duration"`
		School        string            `json:"school"`
		Concentration bool              `json:"concentration"`
		Range         string            `json:"range"`
		Components    []string          `json:"componenets"`
		FlavourText   string            `json:"flavourtext"`
		Classes       string            `json:"classes"`
		SubClasses    string            `json:"subclasses"`
		Source        string            `json:"source"`
		Type          string            `json:"type"`
		AOEShape      string            `json:"aoeshape"`
		AOERadius     int               `json:"aoeradius"`
		SaveAttribute string            `json:"saveattribute"`
		Damage        map[string]string `json:"damage"`
		SaveEffect    string            `json:"saveeffect"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&SpellRequest)

	if jsonParseError != nil {
		http.Error(response, "Unable to Parse JSON", http.StatusBadRequest)
		return
	}

	var spell interface{}

	switch SpellRequest.Type {
	case "AttackBasedRangeAOEAttack":
		{
			spell = models.AttackBasedRangeAOEAttack{
				Spells: models.Spells{
					Name:          SpellRequest.Name,
					Level:         SpellRequest.Level,
					CastingTime:   SpellRequest.CastingTime,
					Duration:      SpellRequest.Duration,
					School:        models.SchoolOfMagic(SpellRequest.School),
					Concentration: SpellRequest.Concentration,
					Range:         SpellRequest.Range,
					Components:    SpellRequest.Components,
					FlavourText:   SpellRequest.FlavourText,
					Classes:       SpellRequest.Classes,
					SubClasses:    SpellRequest.SubClasses,
					SourceName:    SpellRequest.Source,
				},
				AOEShape:      SpellRequest.AOEShape,
				AOERadius:     SpellRequest.AOERadius,
				SaveAttribute: SpellRequest.SaveAttribute,
				Damage:        SpellRequest.Damage,
				SaveEffect:    SpellRequest.SaveEffect,
			}
		}
	}

	if spell == nil {
		http.Error(response, "Unsupported spell type", http.StatusBadRequest)
		return
	}

	insertResult, insertError := database.Spells.InsertOne(context.TODO(), spell)

	if insertError != nil {
		http.Error(response, "Failed to insert spell", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "Spell added successfully",
		"id":      insertResult.InsertedID,
	})
}
