package models

type Backgrounds struct {
	Name                       string   `bson:"name"`
	AbilityScoresUpdated       []string `bson:"abilityscoresupdate"`
	AbilityScoresUpdatedValues []int    `bson:"abilityscorevaluesvalues"`
	SkillProficiencies         []string `bson:"skills"`
	ToolProficiencies          string   `bson:"tools"`
	Languages                  string   `bson:"languages"`
	Feat                       string   `bson:"feat"`
	Gold                       int      `bson:"gold"`
}
