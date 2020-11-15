package events

import (
	"CatLegends/game"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func NewPlayer(cb *tgbotapi.CallbackConfig, chatId int64, msgId int) (tgbotapi.Chattable, bool) {
	p := game.InitPlayer(chatId)

	p.Inventory.Items = []game.Item{
		{
			Name:        "A",
			Emoji:       game.SwordEmoji,
			Quantity:    1,
			Description: "Sword",
			Price:       10,
			Rarity:      game.CommonRarity,
		},
		{
			Name:        "B",
			Emoji:       game.BowEmoji,
			Quantity:    1,
			Description: "Bow",
			Price:       15,
			Rarity:      game.CommonRarity,
		},
	}

	db := utils.GetDB()

	_, ok := game.GetPlayerById(chatId)

	if !ok {
		_, err := db.Players.InsertOne(db.Ctx, p)
		if err != nil {
			log.Error(err)
			cb.Text = ErrorText
			cb.ShowAlert = true
			return nil, false
		}
	}

	cb.Text = "Персонаж створений"
	//editedMsg := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0, 0)})
	editedMsg := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, existingPlayerKeyboard)
	return editedMsg, true
}
