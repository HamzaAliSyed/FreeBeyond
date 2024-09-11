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
	mux.HandleFunc("/api/createitem", handleCreateItem)

}

func handleCreateItem(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	if methodError := utils.OnlyPost(response, request); methodError != nil {
		http.Error(response, methodError.Error(), http.StatusMethodNotAllowed)
		return
	}

	var CreateItemRequest struct {
		Name           string                   `json:"name"`
		Type           string                   `json:"type"`
		Description    string                   `json:"description"`
		Weight         string                   `json:"weight"`
		Cost           string                   `json:"cost"`
		Source         string                   `json:"source"`
		Tags           []string                 `json:"tags"`
		Rarity         models.Rarity            `json:"rarity"`
		Tier           models.Tier              `json:"tier"`
		NeedAttunement bool                     `json:"needattunement"`
		WeaponRangeMin int                      `json:"weaponrangemin"`
		WeaponRangeMax int                      `json:"weaponrangemax"`
		WeaponDamage   map[models.Damage]string `json:"weapondamage"`
	}

	if jsonParseError := json.NewDecoder(request.Body).Decode(&CreateItemRequest); jsonParseError != nil {
		http.Error(response, jsonParseError.Error(), http.StatusBadRequest)
		return
	}

	switch CreateItemRequest.Type {
	case "Weapon":
		{
			weaponItem, weaponItemCreateError := models.CreateNewWeapon(CreateItemRequest.Name, CreateItemRequest.Description, CreateItemRequest.Cost, CreateItemRequest.Weight, CreateItemRequest.Source, CreateItemRequest.Tags, CreateItemRequest.Rarity, CreateItemRequest.Tier, CreateItemRequest.NeedAttunement, CreateItemRequest.WeaponRangeMin, CreateItemRequest.WeaponRangeMax, CreateItemRequest.WeaponDamage)
			if weaponItemCreateError != nil {
				http.Error(response, weaponItemCreateError.Error(), http.StatusInternalServerError)
				return
			}

			itemProperties := weaponItem.GetAllProperties()

			weaponQuery := bson.M{
				"name":               itemProperties["name"].(string),
				"typetags":           itemProperties["typetags"].([]string),
				"rarity":             itemProperties["rarity"].(models.Rarity),
				"tier":               itemProperties["tier"].(models.Tier),
				"requiresAttunement": itemProperties["requiresAttunement"].(bool),
				"description":        itemProperties["description"].(string),
				"cost":               itemProperties["cost"].(string),
				"weight":             itemProperties["weight"].(string),
				"source":             itemProperties["source"].(string),
				"rangemin":           itemProperties["rangemin"].(int),
				"rangemax":           itemProperties["rangemax"].(int),
				"damage":             itemProperties["damage"].(map[models.Damage]string),
			}

			_, insertError := database.Items.InsertOne(context.TODO(), weaponQuery)
			if insertError != nil {
				http.Error(response, insertError.Error(), http.StatusInternalServerError)
				return
			}

			response.WriteHeader(http.StatusCreated)
			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte("New Item that is a weapon has been created"))

		}
	}

}
