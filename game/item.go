package game

const (
	_ int = iota
	CommonRarity
	UncommonRarity
	RareRarity
	EpicRarity
	LegendaryRarity
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
	Rarity      int    `bson:"rarity"`
}
