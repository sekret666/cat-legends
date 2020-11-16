package game

import (
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

type Item struct {
	Name        string `bson:"name"`
	Emoji       string `bson:"emoji"`
	Quantity    int    `bson:"quantity"`
	Description string `bson:"description"`
	Price       int    `bson:"price"`
	Rarity      string `bson:"rarity"`
}

func (i *Item) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%itemName%"] = i.Name
	m["%itemEmoji%"] = i.Emoji
	m["%itemQuantity%"] = strconv.Itoa(i.Quantity)
	m["%itemDescription%"] = i.Description
	m["%itemPrice%"] = strconv.Itoa(i.Price)
	m["%itemRarity%"] = i.Rarity
	return m
}

func (i *Item) ReplaceInString(text string) string {
	for k, v := range i.GetStringMap() {
		text = strings.Replace(text, k, v, -1)
	}
	return text
}
