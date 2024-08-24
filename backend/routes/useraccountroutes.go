package routes

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AccountRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/accounts/createrequest", HandleCreateAccountRequest)
	mux.HandleFunc("/api/accounts/createaccount", HandleCreateAccountUserFormRequest)
	mux.HandleFunc("/api/accounts/usernamematchrequest", HandleUserNameMatchRequest)
	mux.HandleFunc("/api/accounts/signinrequest", HandleSignAccountRequest)
	mux.HandleFunc("/api/accounts/signin", HandleSignInFormRequest)
}

func HandleCreateAccountRequest(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		return
	}

	response.Header().Add("request", "createAccount")
	response.WriteHeader(200)
	response.Write([]byte("Create Account recieved"))
}

func HandleCreateAccountUserFormRequest(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
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

	_, err = database.Users.InsertOne(context.TODO(), userfromform)
	if err != nil {
		http.Error(response, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("User created successfully"))
}

func HandleSignAccountRequest(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		return
	}
	response.Header().Add("request", "signin")
	response.WriteHeader(200)
	response.Write([]byte("Sign in request recieved"))
}

func HandleSignInFormRequest(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
		return
	}
	var sessioncredentials Credentials
	sessionerror := json.NewDecoder(request.Body).Decode(&sessioncredentials)

	if sessionerror != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User

	findError := database.Users.FindOne(context.TODO(), bson.M{"username": sessioncredentials.Username}).Decode(&user)
	if findError != nil {
		http.Error(response, "User not found", http.StatusUnauthorized)
		return
	}

	passwordcompareerror := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(sessioncredentials.Password))

	if passwordcompareerror != nil {
		http.Error(response, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, tokenerror := utils.GenerateJWT(user.Username)
	if tokenerror != nil {
		http.Error(response, "Error generating token", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"token": token,
	})
}

func HandleUserNameMatchRequest(response http.ResponseWriter, request *http.Request) {
	utils.AllowCorsHeaderAndPreflight(response, request)
	methoderror := utils.OnlyPost(response, request)
	if methoderror != nil {
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

	userExists := utils.CheckIfUsernameIsInDatabase(requestData.Username)

	if userExists {
		http.Error(response, "Username already exists", http.StatusConflict)
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
