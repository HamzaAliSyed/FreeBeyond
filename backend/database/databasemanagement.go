package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func ConnectToMongo() {
	mongoClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var ConnectionToMongo error
	mongoClient, ConnectionToMongo = mongo.Connect(context.TODO(), mongoClientOptions)

	if ConnectionToMongo != nil {
		log.Fatal(ConnectionToMongo)
	}

	ConnectionToMongo = mongoClient.Ping(context.TODO(), readpref.Primary())
	if ConnectionToMongo != nil {
		log.Fatal(ConnectionToMongo)
	}

	fmt.Println("Connected to MongoDB!")

}
