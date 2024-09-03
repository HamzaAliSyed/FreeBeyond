package utils

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func FindItemObjectID(itemName string) (primitive.ObjectID, error) {
	var item models.Items
	err := database.Items.FindOne(context.TODO(), bson.M{"name": itemName}).Decode(&item)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return item.ID, nil
}

func FindToolObjectID(toolName string) (primitive.ObjectID, error) {
	var tool struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err := database.ArtisianTools.FindOne(context.TODO(), bson.M{"name": toolName}).Decode(&tool)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, fmt.Errorf("tool not found: %s", toolName)
		}
		return primitive.NilObjectID, err
	}

	return tool.ID, nil
}

func FindClassObjectID(className string) (primitive.ObjectID, error) {
	var class struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err := database.Classes.FindOne(context.TODO(), bson.M{"name": className}).Decode(&class)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, fmt.Errorf("class not found: %s", className)
		}
		return primitive.NilObjectID, err
	}

	return class.ID, nil

}
func ConvertNamesToObjectIDs(collection *mongo.Collection, names []string) ([]primitive.ObjectID, error) {
	var ids []primitive.ObjectID
	for _, name := range names {
		var result struct {
			ID primitive.ObjectID `bson:"_id"`
		}
		err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("invalid name: %s", name)
		}
		ids = append(ids, result.ID)
	}
	return ids, nil
}

func StripBase64Prefix(input string) string {
	commaIndex := strings.Index(input, ",")
	if commaIndex != -1 {
		return input[commaIndex+1:]
	}
	return input
}

func GenerateSavingThrows(abilityScores []models.AbilityScore, profBonus int) []models.SavingThrow {
	savingThrows := make([]models.SavingThrow, 0)

	for _, ability := range abilityScores {
		savingThrow := models.SavingThrow{
			Ability:               ability.AbilityName,
			NumberOfProficiencies: 0, // Default to 0
			HasAdvantage:          false,
			HasDisadvantage:       false,
			Value:                 ability.AbilityModifier,
		}
		savingThrows = append(savingThrows, savingThrow)
	}

	return savingThrows
}

func GetClassFromName(className string) (class *models.Class, fetchError error) {
	class = &models.Class{}
	classFetchError := database.Classes.FindOne(context.TODO(), bson.M{"name": className}).Decode(class)
	if classFetchError != nil {
		if classFetchError == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("class not found: %s", className)
		}
		return nil, classFetchError
	}

	return class, nil
}
