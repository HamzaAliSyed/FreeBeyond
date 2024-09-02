package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/components/getabilitymodifier", getAbilityModifier)
	mux.HandleFunc("/api/components/createsource", handlecreatesource)
	mux.HandleFunc("/api/components/getallsources", getAllSources)
	mux.HandleFunc("/api/components/getallsourcesnames", getAllSourcesNames)
	mux.HandleFunc("/api/components/addspells", handleAddSpells)
	mux.HandleFunc("/api/components/createclass", handleCreateClass)
	mux.HandleFunc("/api/components/createsubclass", handleCreateSubClass)
	mux.HandleFunc("/api/components/createitems", handleCreateNewItems)
	mux.HandleFunc("/api/components/createartisiantools", handleCreateArtisianTools)
	mux.HandleFunc("/api/components/getallitems", getAllItems)
	mux.HandleFunc("/api/components/getallartisiantools", getAllArtisianTools)
	mux.HandleFunc("/api/components/getallclasses", getAllClasses)
	mux.HandleFunc("/api/components/getallsubclasses", getAllSubClasses)
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

func handleCreateClass(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if err := utils.OnlyPost(response, request); err != nil {
		http.Error(response, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var ClassRequest struct {
		Name                   string   `json:"name"`
		HitDie                 string   `json:"hitdie"`
		ArmorProficiency       []string `json:"armorProficiency"`
		WeaponProficiency      []string `json:"weaponProficiency"`
		ToolsProficiency       []string `json:"toolsProficiency"`
		SavingThrowProficiency []string `json:"savingThrowProficiency"`
		SkillsCanChoose        int      `json:"skillsCanChoose"`
		SkillsChoiceList       []string `json:"skillsChoiceList"`
		ToolProficiencies      []string `json:"toolProficiencies"`
		Source                 string   `json:"source"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&ClassRequest)

	if jsonParseError != nil {
		http.Error(response, "Invalid JSON", http.StatusBadRequest)
		return
	}

	toolProficiencyIDs := make([]primitive.ObjectID, 0)
	for _, toolName := range ClassRequest.ToolProficiencies {
		id, err := utils.FindToolObjectID(toolName)
		if err != nil {
			http.Error(response, fmt.Sprintf("Invalid tool: %s", toolName), http.StatusBadRequest)
			return
		}
		toolProficiencyIDs = append(toolProficiencyIDs, id)
	}

	sourceID, err := utils.FindSourceObjectID(ClassRequest.Source)
	if err != nil {
		http.Error(response, fmt.Sprintf("Invalid source: %s", ClassRequest.Source), http.StatusBadRequest)
		return
	}

	newClass := models.Class{
		Name:                   ClassRequest.Name,
		HitDie:                 ClassRequest.HitDie,
		ArmorProficiency:       ClassRequest.ArmorProficiency,
		WeaponProficiency:      ClassRequest.WeaponProficiency,
		ToolsProficiency:       ClassRequest.ToolsProficiency,
		SavingThrowProficiency: ClassRequest.SavingThrowProficiency,
		SkillsCanChoose:        ClassRequest.SkillsCanChoose,
		SkillsChoiceList:       ClassRequest.SkillsChoiceList,
		ToolProficiencies:      toolProficiencyIDs,
		Source:                 sourceID,
		SubClasses:             []primitive.ObjectID{},
	}

	insertResult, insertResultError := database.Classes.InsertOne(context.TODO(), newClass)
	if insertResultError != nil {
		http.Error(response, "Error inserting class", http.StatusInternalServerError)
		return
	}

	newClass.ID = insertResult.InsertedID.(primitive.ObjectID)

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(newClass)
}

func handleCreateSubClass(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if err := utils.OnlyPost(response, request); err != nil {
		http.Error(response, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var SubClassRequest struct {
		Name        string `json:"name"`
		ParentClass string `json:"parentclass"`
		Source      string `json:"source"`
	}

	if err := json.NewDecoder(request.Body).Decode(&SubClassRequest); err != nil {
		http.Error(response, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	parentClassID, err := utils.FindClassObjectID(SubClassRequest.ParentClass)
	if err != nil {
		http.Error(response, fmt.Sprintf("Invalid parent class: %s", err), http.StatusBadRequest)
		return
	}

	sourceID, err := utils.FindSourceObjectID(SubClassRequest.Source)
	if err != nil {
		http.Error(response, fmt.Sprintf("Invalid source: %s", err), http.StatusBadRequest)
		return
	}

	var parentClass models.Class
	err = database.Classes.FindOne(context.TODO(), bson.M{"_id": parentClassID}).Decode(&parentClass)
	if err != nil {
		log.Printf("Error fetching parent class: %v", err)
		http.Error(response, "Error fetching parent class", http.StatusInternalServerError)
		return
	}

	newSubClass := models.SubClasses{
		Name:        SubClassRequest.Name,
		ParentClass: parentClassID,
		Source:      sourceID,
	}

	insertResult, err := database.SubClasses.InsertOne(context.TODO(), newSubClass)
	if err != nil {
		log.Printf("Error inserting subclass: %v", err)
		http.Error(response, "Error inserting subclass", http.StatusInternalServerError)
		return
	}

	newSubClassID := insertResult.InsertedID.(primitive.ObjectID)

	if parentClass.SubClasses == nil {
		parentClass.SubClasses = []primitive.ObjectID{}
	}

	parentClass.SubClasses = append(parentClass.SubClasses, newSubClassID)

	updateResult, err := database.Classes.UpdateOne(
		context.TODO(),
		bson.M{"_id": parentClassID},
		bson.M{"$set": bson.M{"subclasses": parentClass.SubClasses}},
	)

	if err != nil {
		log.Printf("Error updating parent class: %v", err)
		http.Error(response, fmt.Sprintf("Error updating parent class: %v", err), http.StatusInternalServerError)
		return
	}

	if updateResult.MatchedCount == 0 {
		log.Printf("No matching document found for parent class ID: %v", parentClassID)
		http.Error(response, "Parent class not found", http.StatusNotFound)
		return
	}

	if updateResult.ModifiedCount == 0 {
		log.Printf("Document matched but not modified. This might happen if the subclass was already in the array.")
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "Subclass created successfully",
		"id":      newSubClassID,
	})
}

func handleCreateNewItems(response http.ResponseWriter, request *http.Request) {

	utils.AllowCorsHeaderAndPreflight(response, request)
	if err := utils.OnlyPost(response, request); err != nil {
		http.Error(response, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var ItemCreateRequest struct {
		Name               string                    `json:"name"`
		TypeTags           []string                  `json:"typeTags"`
		ItemType           string                    `json:"itemType"`
		RequiresAttunement bool                      `json:"requiresAttunement"`
		Cost               string                    `json:"cost"`
		Weight             string                    `json:"weight"`
		FlavourText        []models.TextBasedAbility `json:"flavourText"`
		Source             string                    `json:"source"`
	}

	jsonParseError := json.NewDecoder(request.Body).Decode(&ItemCreateRequest)
	if jsonParseError != nil {
		http.Error(response, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	itemType := models.ItemType(ItemCreateRequest.ItemType)
	if !utils.IsValidItemType(itemType) {
		http.Error(response, "Invalid ItemType", http.StatusBadRequest)
		return
	}

	sourceObjectID, err := utils.FindSourceObjectID(ItemCreateRequest.Source)
	if err != nil {
		http.Error(response, "Invalid Source", http.StatusBadRequest)
		return
	}

	newItem := models.Items{
		Name:               ItemCreateRequest.Name,
		TypeTags:           ItemCreateRequest.TypeTags,
		ItemType:           itemType,
		RequiresAttunement: ItemCreateRequest.RequiresAttunement,
		Cost:               ItemCreateRequest.Cost,
		Weight:             ItemCreateRequest.Weight,
		FlavourText:        ItemCreateRequest.FlavourText,
		Source:             sourceObjectID,
	}

	insertResult, insertResultError := database.Items.InsertOne(context.TODO(), newItem)
	if insertResultError != nil {
		http.Error(response, "Error inserting item into database", http.StatusInternalServerError)
		return
	}

	insertedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		http.Error(response, "Error retrieving inserted item ID", http.StatusInternalServerError)
		return
	}

	responseItem := struct {
		ID   primitive.ObjectID `json:"id"`
		Name string             `json:"name"`
	}{
		ID:   insertedID,
		Name: newItem.Name,
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(responseItem)
}

func handleCreateArtisianTools(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if err := utils.OnlyPost(response, request); err != nil {
		http.Error(response, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var artisianToolsRequest struct {
		Name        string                    `json:"name"`
		Description string                    `json:"description"`
		Items       []string                  `json:"items"`
		FlavourText []models.TextBasedAbility `json:"flavourText"`
		DCTable     map[string]int            `json:"dcTable"`
	}

	if err := json.NewDecoder(request.Body).Decode(&artisianToolsRequest); err != nil {
		http.Error(response, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	var itemIDs []primitive.ObjectID
	for _, itemName := range artisianToolsRequest.Items {
		id, err := utils.FindItemObjectID(itemName)
		if err != nil {
			http.Error(response, fmt.Sprintf("Invalid item: %s", itemName), http.StatusBadRequest)
			return
		}
		itemIDs = append(itemIDs, id)
	}

	newArtisianTools := models.ArtisianTools{
		Name:        artisianToolsRequest.Name,
		Description: artisianToolsRequest.Description,
		Items:       itemIDs,
		FlavourText: artisianToolsRequest.FlavourText,
		DCTable:     artisianToolsRequest.DCTable,
	}

	insertResult, err := database.ArtisianTools.InsertOne(context.TODO(), newArtisianTools)
	if err != nil {
		http.Error(response, "Error inserting artisian tools", http.StatusInternalServerError)
		return
	}

	insertedID := insertResult.InsertedID.(primitive.ObjectID)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(map[string]interface{}{
		"id":   insertedID,
		"name": newArtisianTools.Name,
	})
}

func getAllItems(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyGet(response, request)

	if methodError != nil {
		http.Error(response, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, cursorError := database.Items.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"name": 1}))

	if cursorError != nil {
		http.Error(response, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var itemNames []string

	for cursor.Next(context.TODO()) {
		var result struct {
			Name string `bson:"name"`
		}
		if cursorError := cursor.Decode(&result); cursorError != nil {
			http.Error(response, "Failed to decode data", http.StatusInternalServerError)
			return
		}
		itemNames = append(itemNames, result.Name)
	}

	if internalServerError := cursor.Err(); internalServerError != nil {
		http.Error(response, "Cursor error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(itemNames)
}

func getAllArtisianTools(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyGet(response, request)

	if methodError != nil {
		http.Error(response, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, cursorError := database.ArtisianTools.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"name": 1}))

	if cursorError != nil {
		http.Error(response, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var toolNames []string

	for cursor.Next(context.TODO()) {
		var result struct {
			Name string `bson:"name"`
		}
		if cursorError := cursor.Decode(&result); cursorError != nil {
			http.Error(response, "Failed to decode data", http.StatusInternalServerError)
			return
		}
		toolNames = append(toolNames, result.Name)
	}

	if internalServerError := cursor.Err(); internalServerError != nil {
		http.Error(response, "Cursor error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(toolNames)
}

func getAllClasses(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyGet(response, request)
	if methodError != nil {
		http.Error(response, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, cursorError := database.Classes.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"name": 1}))

	if cursorError != nil {
		http.Error(response, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var classNames []string

	for cursor.Next(context.TODO()) {
		var class struct {
			Name string `bson:"name"`
		}

		if cursorError := cursor.Decode(&class); cursorError != nil {
			http.Error(response, "Failed to decode data", http.StatusInternalServerError)
			return
		}

		classNames = append(classNames, class.Name)
	}

	if internalServerError := cursor.Err(); internalServerError != nil {
		http.Error(response, "Cursor error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&classNames)
}

func getAllSubClasses(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methodError := utils.OnlyGet(response, request)
	if methodError != nil {
		http.Error(response, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, cursorError := database.SubClasses.Find(context.TODO(), bson.M{}, options.Find().SetProjection(bson.M{"name": 1}))

	if cursorError != nil {
		http.Error(response, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var subClassNames []string

	for cursor.Next(context.TODO()) {
		var subClass struct {
			Name string `bson:"name"`
		}

		if cursorError := cursor.Decode(&subClass); cursorError != nil {
			http.Error(response, "Failed to decode data", http.StatusInternalServerError)
			return
		}

		subClassNames = append(subClassNames, subClass.Name)
	}

	if internalServerError := cursor.Err(); internalServerError != nil {
		http.Error(response, "Cursor error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&subClassNames)
}
