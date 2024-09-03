package main

import (
	"backend/database"
	"backend/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	environmentPath := "../.env"
	environmenterror := godotenv.Load(environmentPath)
	if environmenterror != nil {
		log.Fatalf("Error loading .env file")
	}

	const port = "2712"
	database.ConnectToMongo()
	backend := http.NewServeMux()
	routes.CharacterRoutes(backend)
	log.Fatal(http.ListenAndServe(":"+port, backend))
}
