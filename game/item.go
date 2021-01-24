package game

import (
	"CatLegends/game/items"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
)

const (
	CommonRarity    = "Звичайний"
	UncommonRarity  = "Незвичайний"
	RareRarity      = "Рідкісний"
	EpicRarity      = "Епічний"
	LegendaryRarity = "Легендарний"
)

const (
	SwordEmoji       = "🗡"
	BowEmoji         = "🏹"
	ClothingEmoji    = "👕"
	JewelryEmoji     = "📿"
	AccessoriesEmoji = "🌂"
)

type ItemDetails interface {
	Info(pattern string) string
	DefaultPattern() string
}

type Item struct {
	Name            string      `bson:"name"`
	Emoji           string      `bson:"emoji"`
	Quantity        int         `bson:"quantity"`
	Description     string      `bson:"description"`
	Price           int         `bson:"price"`
	Rarity          string      `bson:"rarity"`
	ItemDetailsType string      `bson:"itemDetailsType"`
	ItemDetails     ItemDetails `bson:"itemDetails"`
}

func (i *Item) UnmarshalBSON(bytes []byte) error {
	var raw bson.Raw
	if err := bson.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	var itemDetailsType string
	if err := raw.Lookup("itemDetailsType").Unmarshal(&itemDetailsType); err != nil {
		return err
	}

	switch itemDetailsType {
	case items.WeaponType:
		var w items.Weapon
		if err := raw.Lookup("itemDetails").Unmarshal(&w); err != nil {
			return err
		}
		i.ItemDetails = &w
	case items.ClothingType:
		var c items.Clothing
		if err := raw.Lookup("itemDetails").Unmarshal(&c); err != nil {
			return err
		}
		i.ItemDetails = &c
	case items.AccessoriesType:
		var a items.Accessory
		if err := raw.Lookup("itemDetails").Unmarshal(&a); err != nil {
			return err
		}
		i.ItemDetails = &a
	default:
		return fmt.Errorf("unkown item details type: %s", itemDetailsType)
	}

	if err := raw.Lookup("name").Unmarshal(&i.Name); err != nil {
		return err
	}

	if err := raw.Lookup("emoji").Unmarshal(&i.Emoji); err != nil {
		return err
	}

	if err := raw.Lookup("quantity").Unmarshal(&i.Quantity); err != nil {
		return err
	}

	if err := raw.Lookup("description").Unmarshal(&i.Description); err != nil {
		return err
	}

	if err := raw.Lookup("price").Unmarshal(&i.Price); err != nil {
		return err
	}

	if err := raw.Lookup("rarity").Unmarshal(&i.Rarity); err != nil {
		return err
	}

	return nil
}

func (i *Item) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%itemName%"] = i.Name
	m["%itemEmoji%"] = i.Emoji
	m["%itemQuantity%"] = strconv.Itoa(i.Quantity)
	m["%itemDescription%"] = i.Description
	m["%itemPrice%"] = strconv.Itoa(i.Price)
	m["%itemRarity%"] = i.Rarity
	m["%itemTypeInfo%"] = i.ItemDetails.Info(i.ItemDetails.DefaultPattern())
	return m
}

func (i *Item) ReplaceInString(text string) string {
	for k, v := range i.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
