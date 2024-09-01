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
var Sources *mongo.Collection
var Spells *mongo.Collection
var Classes *mongo.Collection

func ConnectToMongo() {
	MongoClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var ConnectionToMongo error
	mongoClient, ConnectionToMongo = mongo.Connect(context.TODO(), MongoClientOptions)

	if ConnectionToMongo != nil {
		log.Fatal(ConnectionToMongo)
	}

	ConnectionToMongo = mongoClient.Ping(context.TODO(), readpref.Primary())
	if ConnectionToMongo != nil {
		log.Fatal(ConnectionToMongo)
	}

	fmt.Println("Connected to MongoDB!")

	Sources = mongoClient.Database("DND").Collection("Sources")
	Spells = mongoClient.Database("DND").Collection("Spells")
	Classes = mongoClient.Database("DND").Collection("Classes")

}
