package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	loadEnvironmentError := godotenv.Load()
	if loadEnvironmentError != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	CERT_PATH := os.Getenv("CERT_PATH")
	KEY_PATH := os.Getenv("KEY_PATH")

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("HTTPS Test Page"))
	})

	httpsError := http.ListenAndServeTLS(":"+PORT, CERT_PATH, KEY_PATH, nil)
	if httpsError != nil {
		log.Fatal(httpsError)
	}
}
