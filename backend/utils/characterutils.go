package utils

import (
	"backend/database"
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsValidItemType(itemType models.ItemType) bool {
	validTypes := []models.ItemType{
		models.Mundane, models.Common, models.Uncommon,
		models.Rare, models.VeryRare, models.Legendary, models.Artifact,
	}
	for _, t := range validTypes {
		if t == itemType {
			return true
		}
	}
	return false
}

func FindSourceObjectID(sourceName string) (primitive.ObjectID, error) {
	var source models.Source
	err := database.Sources.FindOne(context.TODO(), bson.M{"name": sourceName}).Decode(&source)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return source.ID, nil
}
