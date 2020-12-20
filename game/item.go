package game

import (
	"CatLegends/game/items"
	"errors"
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

type ItemTypeRaw struct {
	Type string   `bson:"type"`
	Raw  bson.Raw `bson:"raw"`
}

func (i ItemTypeRaw) ItemTypeToRaw(it ItemType) (err error) {
	i.Raw, err = bson.Marshal(it)
	return err
}

func (i ItemTypeRaw) RawToItemType() (it ItemType, err error) {
	switch i.Type {
	case "weapon":
		var w items.Weapon
		err = bson.Unmarshal(i.Raw, &w)
		it = &w
	default:
		err = errors.New("invalid item type")
	}
	return it, err
}

type ItemType interface {
	Info(pattern string) string
	DefaultPattern() string
}

type Item struct {
	Name        string      `bson:"name"`
	Emoji       string      `bson:"emoji"`
	Quantity    int         `bson:"quantity"`
	Description string      `bson:"description"`
	Price       int         `bson:"price"`
	Rarity      string      `bson:"rarity"`
	ItemType    ItemType    `bson:"-"`
	ItemTypeRaw ItemTypeRaw `bson:"itemTypeRaw"`
}

func (i *Item) GetStringMap() map[string]string {
	m := make(map[string]string)
	m["%itemName%"] = i.Name
	m["%itemEmoji%"] = i.Emoji
	m["%itemQuantity%"] = strconv.Itoa(i.Quantity)
	m["%itemDescription%"] = i.Description
	m["%itemPrice%"] = strconv.Itoa(i.Price)
	m["%itemRarity%"] = i.Rarity
	m["%itemTypeInfo%"] = i.ItemType.Info(i.ItemType.DefaultPattern())
	return m
}

func (i *Item) ReplaceInString(text string) string {
	for k, v := range i.GetStringMap() {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
