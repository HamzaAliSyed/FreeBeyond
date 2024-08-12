package main

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleCreateAccountRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if request.Method == "POST" {
		response.Header().Add("request", "createAccount")
		response.WriteHeader(200)
		response.Write([]byte("Create Account recieved"))
	} else {
		response.Write([]byte("Method should be POST only"))
	}
}

func HandleCreateAccountUserFormRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if request.Method != "POST" {
		http.Error(response, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var userfromform models.User
	err := json.NewDecoder(request.Body).Decode(&userfromform)
	if err != nil {
		http.Error(response, "Error decoding form data", http.StatusBadRequest)
		return
	}

	_, err = users.InsertOne(context.TODO(), userfromform)
	if err != nil {
		http.Error(response, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	// Send HTTP 201 Created status if the user is successfully inserted
	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("User created successfully"))

}
