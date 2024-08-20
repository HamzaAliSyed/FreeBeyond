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
var Users *mongo.Collection
var Characters *mongo.Collection
var Skills *mongo.Collection
var Feats *mongo.Collection
var Sources *mongo.Collection
var Races *mongo.Collection
var Classes *mongo.Collection

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

	Users = mongoClient.Database("dnd").Collection("users")
	Characters = mongoClient.Database("dnd").Collection("characters")
	Skills = mongoClient.Database("dnd").Collection("skills")
	Feats = mongoClient.Database("dnd").Collection("feats")
	Sources = mongoClient.Database("dnd").Collection("sources")
	Races = mongoClient.Database("dnd").Collection("races")
	Classes = mongoClient.Database("dnd").Collection("classes")

}
