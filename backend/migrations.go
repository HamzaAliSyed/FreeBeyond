package main

import (
	"backend/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Migration struct {
	CollectionName string
	Update         bson.M
	Filter         bson.M
}

var migrations = []Migration{
	{
		CollectionName: "classes",
		Filter:         bson.M{"subclasses": bson.M{"$exists": false}},
		Update:         bson.M{"$set": bson.M{"subclasses": []interface{}{}}},
	},
}

func fixClassSubclasses(db *mongo.Database) error {
	cursor, err := db.Collection("classes").Find(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("error fetching classes: %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var class models.Class
		if err := cursor.Decode(&class); err != nil {
			return fmt.Errorf("error decoding class: %v", err)
		}

		if class.SubClasses == nil {
			class.SubClasses = []primitive.ObjectID{}
			_, err := db.Collection("classes").UpdateOne(
				context.TODO(),
				bson.M{"_id": class.ID},
				bson.M{"$set": bson.M{"subclasses": class.SubClasses}},
			)
			if err != nil {
				return fmt.Errorf("error updating class %s: %v", class.ID, err)
			}
		}
	}

	return nil
}

func runMigrations(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, migration := range migrations {
		collection := db.Collection(migration.CollectionName)
		log.Printf("Running migration on collection: %s", migration.CollectionName)
		result, err := collection.UpdateMany(ctx, migration.Filter, migration.Update)
		if err != nil {
			log.Printf("Error running migration on collection %s: %v", migration.CollectionName, err)
			return err
		}
		log.Printf("Migration completed. Modified %d document(s)", result.ModifiedCount)
	}

	if err := fixClassSubclasses(db); err != nil {
		log.Printf("Error fixing class subclasses: %v", err)
		return err
	}

	return nil
}
