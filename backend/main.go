package main

import (
	"backend/database"
	"backend/routes"
	"context"
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
	defer database.MongoClient.Disconnect(context.Background())
	DnDDatabase := database.MongoClient.Database("DND")
	log.Println("Starting database migrations...")
	if migrationError := runMigrations(DnDDatabase); migrationError != nil {
		log.Fatalf("Failed to run migrations: %v", migrationError)
	}
	log.Println("Database migrations completed successfully")
	backend := http.NewServeMux()
	routes.HandleComponentRoutes(backend)
	routes.HandleCharacterRoutes(backend)
	log.Fatal(http.ListenAndServe(":"+port, backend))
}
