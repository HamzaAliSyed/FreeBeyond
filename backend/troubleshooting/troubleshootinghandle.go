package troubleshooting

import (
	"backend/database"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func TroubleshootingRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/troubleshooting/collection", CollectionTest)
}

func CollectionTest(response http.ResponseWriter, request *http.Request) {

	cursor, err := database.Characters.Find(context.TODO(), bson.D{{}})
	if err != nil {
		http.Error(response, fmt.Sprintf("Failed to access collection: %v", err), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			http.Error(response, fmt.Sprintf("Failed to decode document: %v", err), http.StatusInternalServerError)
			return
		}
		// Print the _id field of the document
		if id, ok := result["_id"]; ok {
			fmt.Fprintf(response, "_id: %v\n", id)
		} else {
			fmt.Fprintf(response, "Document does not have _id field\n")
		}
	}

	if err := cursor.Err(); err != nil {
		http.Error(response, fmt.Sprintf("Cursor error: %v", err), http.StatusInternalServerError)
	}

}
