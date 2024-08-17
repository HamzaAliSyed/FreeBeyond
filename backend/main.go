package main

import (
	"backend/database"
	"backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func genericOk(reponse http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)

	fileServer := http.FileServer(http.Dir("."))
	fileServer.ServeHTTP(reponse, request)

	reponse.WriteHeader(http.StatusOK)
}

func main() {
	environmenterror := godotenv.Load()
	if environmenterror != nil {
		log.Fatalf("Error loading .env file")
	}
	database.ConnectToMongo()
	const port = "2712"
	backend := http.NewServeMux()
	routes.FeatRoutes(backend)
	routes.AccountRoutes(backend)

	backend.HandleFunc("/", genericOk)
	log.Fatal(http.ListenAndServe(":"+port, backend))
}
