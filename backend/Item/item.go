package item

import "errors"

type CreateParams interface{}

type Item interface {
	Create(params CreateParams) error
	Print()
	GetItemName() string
	GetItemTags() []string
}

type Rarity string

const (
	Mundane   Rarity = "Mundane"
	Common    Rarity = "Common"
	Uncommon  Rarity = "Uncommon"
	Rare      Rarity = "Rare"
	VeryRare  Rarity = "Very Rare"
	Legendary Rarity = "Legendary"
	Artifact  Rarity = "Artifact"
)

func IsValidRarity(r Rarity) bool {
	switch r {
	case Mundane, Common, Uncommon, Rare, VeryRare, Legendary, Artifact:
		return true
	}
	return false
}

type BasicParameter struct {
	name        string   `bson:"name"`
	typeTags    []string `bson:"typetags"`
	description string   `bson:"description"`
	rarity      Rarity   `bson:"rarity"`
	cost        string   `bson:"cost"`
	weight      string   `bson:"weight"`
	source      string   `bson:"source"`
}

func CreateItemFactory(itemType string, params CreateParams) (Item, error) {
	switch itemType {
	case "weapon":
		{
			weapon := &Weapon{}
			weaponCreateError := weapon.Create(params)
			return weapon, weaponCreateError
		}
	}

	return nil, errors.New("invalid Item Type")
}
