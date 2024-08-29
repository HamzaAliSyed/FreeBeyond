package main

import (
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
	backend := http.NewServeMux()
	routes.HandleComponentRoutes(backend)
	log.Fatal(http.ListenAndServe(":"+port, backend))
}
