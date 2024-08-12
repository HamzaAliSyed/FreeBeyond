package main

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
		return
	}
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

	if userfromform.FirstName == "" || userfromform.LastName == "" || userfromform.Email == "" || userfromform.Password == "" || userfromform.Username == "" {
		http.Error(response, "Missing required fields", http.StatusBadRequest)
		return
	}

	hashedPassword, PassError := bcrypt.GenerateFromPassword([]byte(userfromform.Password), bcrypt.MinCost)
	if PassError != nil {
		http.Error(response, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userfromform.Password = string(hashedPassword)

	_, err = users.InsertOne(context.TODO(), userfromform)
	if err != nil {
		http.Error(response, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	// Send HTTP 201 Created status if the user is successfully inserted
	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("User created successfully"))

}

func HandleUserNameMatchRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
		return
	}
	if request.Method != "POST" {
		http.Error(response, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Username string `json:"Username"`
	}

	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		http.Error(response, "Error decoding request", http.StatusBadRequest)
		return
	}

	userExists := CheckIfUsernameIsInDatabase(requestData.Username)

	if userExists {
		http.Error(response, "Username already exists", http.StatusConflict)
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

func CheckIfUsernameIsInDatabase(username string) bool {
	filter := bson.M{"username": username}
	var user models.User

	err := users.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false // Username does not exist
		}
		log.Printf("Error checking username in database: %v", err)
		return false
	}

	return true
}
