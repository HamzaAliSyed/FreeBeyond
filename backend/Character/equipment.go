package character

type Equipment struct {
	head     string `bson:"head"`
	mainhand string `bson:"mainhand"`
	offhand  string `bson:"sidehand"`
	wrist    string `bson:"wrist"`
	torso    string `bson:"torso"`
	back     string `bson:"back"`
	legs     string `bson:"legs"`
}
