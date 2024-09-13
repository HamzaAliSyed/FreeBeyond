package main

import (
	"backend/database"
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

	log.Fatal(http.ListenAndServe(":"+port, backend))
}
