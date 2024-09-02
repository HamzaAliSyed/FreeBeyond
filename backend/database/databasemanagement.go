package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var Characters *mongo.Collection
var Sources *mongo.Collection
var Spells *mongo.Collection
var Classes *mongo.Collection
var Items *mongo.Collection
var ArtisianTools *mongo.Collection
var SubClasses *mongo.Collection
var Races *mongo.Collection

func ConnectToMongo() {
	MongoClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var ConnectionToMongo error
	MongoClient, ConnectionToMongo = mongo.Connect(context.TODO(), MongoClientOptions)

	if ConnectionToMongo != nil {
		log.Fatal(ConnectionToMongo)
	}

	ConnectionToMongo = MongoClient.Ping(context.TODO(), readpref.Primary())
	if ConnectionToMongo != nil {
		log.Fatal(ConnectionToMongo)
	}

	fmt.Println("Connected to MongoDB!")

	Sources = MongoClient.Database("DND").Collection("Sources")
	Spells = MongoClient.Database("DND").Collection("Spells")
	Classes = MongoClient.Database("DND").Collection("Classes")
	Items = MongoClient.Database("DND").Collection("Items")
	ArtisianTools = MongoClient.Database("DND").Collection("ArtisianTools")
	SubClasses = MongoClient.Database("DND").Collection("SubClasses")
	Characters = MongoClient.Database("DND").Collection("Characters")
	Races = MongoClient.Database("DND").Collection("Races")
}
