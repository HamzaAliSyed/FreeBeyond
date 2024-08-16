package utils

import (
	"backend/models"
	"context"
	"fmt"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AllowCorsHeaderAndPreflight(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
		return
	}

}

func OnlyPost(response http.ResponseWriter, request *http.Request) error {
	if request.Method != http.MethodPost {
		http.Error(response, "Only POST method allowed on the end point", http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	return nil
}

func RetrieveCharacter(characterid string, db *mongo.Collection) (*models.Character, error) {

	objectID, objectIDError := primitive.ObjectIDFromHex(characterid)

	if objectIDError != nil {
		return nil, fmt.Errorf("invalid ID format")
	}

	filter := bson.M{"_id": objectID}

	var character models.Character

	charactererror := db.FindOne(context.TODO(), filter).Decode(&character)

	if charactererror != nil {
		return nil, fmt.Errorf("character not found")
	}

	return &character, nil

}

func ModifierCalculator(statvalue int) int {
	statmodifier := (statvalue - 10) / 2
	return statmodifier
}

func InitialSavingThrowsGenerator(character *models.Character) {
	fmt.Println("Generating Initial Saving Throws")
	modifiers := character.Modifiers

	var savingThrows []models.SavingThrow

	valueOfModifier := reflect.ValueOf(modifiers)
	typeOfModifiers := valueOfModifier.Type()

	for i := 0; i < valueOfModifier.NumField(); i++ {
		fieldValue := valueOfModifier.Field(i)
		fieldName := typeOfModifiers.Field(i).Name
		attributeName := fieldName[:len(fieldName)-8]

		savingThrow := models.SavingThrow{
			ID:                    primitive.NewObjectID(),
			Attribute:             attributeName,
			AttributeModifier:     int(fieldValue.Int()),
			SavingThrowValue:      int(fieldValue.Int()),
			NumberOfProficiencies: 0,
		}

		savingThrows = append(savingThrows, savingThrow)
	}

	character.SavingThrow = savingThrows

}

func InitializeSkillsArray(character *models.Character, skills *mongo.Collection) {
	fmt.Println("Generating basic skill array for the character")
	filter := bson.M{}

	skillcursor, skillretrieveerror := skills.Find(context.TODO(), filter)

	if skillretrieveerror != nil {
		fmt.Println(fmt.Errorf("error in retrieving skills: %s", skillretrieveerror))
		return
	}

	var SkillList models.Skills

	for skillcursor.Next(context.TODO()) {
		var skill models.Skill
		if err := skillcursor.Decode(&skill); err != nil {
			fmt.Println("Error decoding skill:", err)
			continue
		}

		switch skill.AssociatedAttribute {
		case "Strength":
			skill.AssociatedAttributeValue = character.Modifiers.StrengthModifier
		case "Dexterity":
			skill.AssociatedAttributeValue = character.Modifiers.DexterityModifier
		case "Constitution":
			skill.AssociatedAttributeValue = character.Modifiers.ConstitutionModifier
		case "Intelligence":
			skill.AssociatedAttributeValue = character.Modifiers.IntelligenceModifier
		case "Wisdom":
			skill.AssociatedAttributeValue = character.Modifiers.WisdomModifier
		case "Charisma":
			skill.AssociatedAttributeValue = character.Modifiers.CharismaModifier
		}

		skill.FinalSkillValue = skill.AssociatedAttributeValue + int(skill.NumberOfProficiencies*float64(skill.ProficiencyBonus)) + skill.AdditionalBoostValue

		SkillList.SkillList = append(SkillList.SkillList, skill)
	}

	character.Skills = SkillList
}
