package events

import (
	"CatLegends/game"
	"CatLegends/game/items"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func NewPlayer(cb *tgbotapi.CallbackConfig, chatId int64, msgId int, bot *tgbotapi.BotAPI) {
	p := game.InitPlayer(chatId)

	p.Inventory.Items = append(p.Inventory.Items, game.Item{
		Name:            "Test Sword",
		Emoji:           game.SwordEmoji,
		Quantity:        1,
		Description:     "Test sword item description",
		Price:           15,
		Rarity:          game.UncommonRarity,
		ItemDetailsType: items.WeaponType,
		ItemDetails: &items.Weapon{
			Damage:    10,
			OneHanded: false,
		},
	})

	p.Inventory.Items = append(p.Inventory.Items, game.Item{
		Name:            "Test Clothing",
		Emoji:           game.ClothingEmoji,
		Quantity:        1,
		Description:     "Test clothing item description",
		Price:           10,
		Rarity:          game.RareRarity,
		ItemDetailsType: items.ClothingType,
		ItemDetails: &items.Clothing{
			Defence: 5,
		},
	})

	p.Inventory.Items = append(p.Inventory.Items, game.Item{
		Name:            "Test Accessory",
		Emoji:           game.AccessoriesEmoji,
		Quantity:        1,
		Description:     "Test accessory item description",
		Price:           20,
		Rarity:          game.CommonRarity,
		ItemDetailsType: items.AccessoriesType,
		ItemDetails: &items.Accessory{
			Effect: "Test effect",
		},
	})

	db := utils.GetDB()

	_, ok := game.GetPlayerById(chatId)

	if !ok {
		_, err := db.Players.InsertOne(db.Ctx, p)
		if err != nil {
			log.Error(err)
			cb.Text = ErrorText
			cb.ShowAlert = true
			return
		}
	}

	cb.Text = "Персонаж створений"
	msgEdit := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, existingPlayerKeyboard)
	if _, err := bot.Send(msgEdit); err != nil {
		log.Error(err)
	}
}
