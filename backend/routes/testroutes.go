package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/database"
	"backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func HandleTestRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/test/getallclassfeatures", handleGetAllClassFeatures)
}

func handleGetAllClassFeatures(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, err := database.ClassFeatures.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error retrieving class features", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var features []models.ClassFeature
	if err = cursor.All(context.TODO(), &features); err != nil {
		http.Error(w, "Error decoding class features", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(features)
}
