package utils

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AllowCorsHeaderAndPreflight(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
		return
	}

}

func OnlyPost(response http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodPost {
		http.Error(response, "Only POST method allowed on the end point", http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	return nil
}

func GenerateJWT(Username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 255)

	claims := &jwt.RegisteredClaims{
		Subject:   Username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	jwtKey := []byte(os.Getenv("JWTKEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckIfUsernameIsInDatabase(username string) bool {
	filter := bson.M{"username": username}
	var user models.User

	err := database.Users.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Printf("Error checking username in database: %v", err)
		return false
	}

	return true
}
