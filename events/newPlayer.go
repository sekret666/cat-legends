package events

import (
	"CatLegends/game"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func NewPlayer(cb *tgbotapi.CallbackConfig, chatId int64, msgId int) (tgbotapi.Chattable, bool) {
	p := game.InitPlayer(chatId)

	db := utils.GetDB()
	_, err := db.Players.InsertOne(db.Ctx, p)
	if err != nil {
		log.Error(err)
		cb.Text = ErrorText
		cb.ShowAlert = true
		return nil, false
	}

	cb.Text = "Персонаж створений"
	//editedMsg := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0, 0)})
	editedMsg := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, existingPlayerKeyboard)
	return editedMsg, true
}
