package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"
)

func HandleSpellRoute(mux *http.ServeMux) {
	mux.HandleFunc("/api/spell/create", handleCreateASpell)
}

func handleCreateASpell(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methodError := utils.OnlyPost(response, request); methodError != nil {
		http.Error(response, methodError.Error(), http.StatusMethodNotAllowed)
		return
	}

	var RequestCreateSpellStruct struct {
		Name                  string            `json:"name"`
		FlavourText           string            `json:"flavourtext"`
		Level                 int               `json:"level"`
		ActionTime            string            `json:"actiontime"`
		CastingTime           string            `json:"castingtime"`
		Components            []string          `json:"components"`
		School                string            `json:"School"`
		RequiresConcentration bool              `json:"requiresconcentration"`
		TypeOfSpell           string            `json:"typeofspell"`
		Range                 string            `json:"range"`
		AccessTo              map[string]string `json:"acessto"`
		Source                string            `json:"source"`
		Damage                map[string]string `json:"damage"`
		Shape                 string            `json:"shape"`
		ShapeSize             string            `json:"shapesize"`
		SaveStat              string            `json:"savestat"`
		SaveEffect            string            `json:"saveeffect"`
		StatsModified         []string          `json:"statsmodified"`
		NumberOfCreature      string            `json:"numberofcreature"`
	}

	if jsonParseError := json.NewDecoder(request.Body).Decode(&RequestCreateSpellStruct); jsonParseError != nil {
		http.Error(response, jsonParseError.Error(), http.StatusBadRequest)
		return
	}

	if !isValidSchool(RequestCreateSpellStruct.School) {
		http.Error(response, "Invalid school", http.StatusBadRequest)
		return
	}

	if !isValidTypeOfSpell(RequestCreateSpellStruct.TypeOfSpell) {
		http.Error(response, "Invalid type of spell", http.StatusBadRequest)
		return
	}

	newSpell := models.Spell{
		Name:                  RequestCreateSpellStruct.Name,
		FlavourText:           RequestCreateSpellStruct.FlavourText,
		Level:                 RequestCreateSpellStruct.Level,
		ActionTime:            RequestCreateSpellStruct.ActionTime,
		CastingTime:           RequestCreateSpellStruct.CastingTime,
		Components:            RequestCreateSpellStruct.Components,
		School:                models.School(RequestCreateSpellStruct.School),
		RequiresConcentration: RequestCreateSpellStruct.RequiresConcentration,
		TypeOfSpell:           models.TypeOfSpell(RequestCreateSpellStruct.TypeOfSpell),
		Range:                 RequestCreateSpellStruct.Range,
		AccessTo:              RequestCreateSpellStruct.AccessTo,
		Source:                RequestCreateSpellStruct.Source,
		Shape:                 RequestCreateSpellStruct.Shape,
		ShapeSize:             RequestCreateSpellStruct.ShapeSize,
		Damage:                RequestCreateSpellStruct.Damage,
		SaveStat:              RequestCreateSpellStruct.SaveStat,
		SaveEffect:            RequestCreateSpellStruct.SaveEffect,
		StatsModified:         RequestCreateSpellStruct.StatsModified,
		NumberofCreature:      RequestCreateSpellStruct.NumberOfCreature,
	}

	insertResult, insertError := database.Spells.InsertOne(context.Background(), newSpell)
	if insertError != nil {
		http.Error(response, "Failed to insert spell into database: "+insertError.Error(), http.StatusInternalServerError)
		return
	}

	if insertResult.InsertedID == nil {
		http.Error(response, "Failed to insert spell: No ID returned", http.StatusInternalServerError)
		return
	}

	successResponse := struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		ID      interface{} `json:"id"`
	}{
		Status:  "success",
		Message: "Spell inserted successfully",
		ID:      insertResult.InsertedID,
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	jsonEncodeError := json.NewEncoder(response).Encode(successResponse)
	if jsonEncodeError != nil {
		http.Error(response, "Failed to encode response: "+jsonEncodeError.Error(), http.StatusInternalServerError)
		return
	}

}

func isValidSchool(s string) bool {
	validSchools := []models.School{
		models.Abjuration,
		models.Conjuration,
		models.Divination,
		models.Enchantment,
		models.Evocation,
		models.Illusion,
		models.Necromancy,
		models.Transmutation,
	}

	for _, validSchool := range validSchools {
		if string(validSchool) == s {
			return true
		}
	}
	return false
}

func isValidTypeOfSpell(t string) bool {
	validTypes := []models.TypeOfSpell{
		models.Generic,
		models.SaveBased,
		models.RollToAttack,
		models.Modifier,
	}

	for _, validType := range validTypes {
		if string(validType) == t {
			return true
		}
	}
	return false
}
