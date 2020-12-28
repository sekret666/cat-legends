package events

import (
	"CatLegends/game"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const startMessage = `
Привіт %name%!
Вітаю тебе у текстовій RPG грі
Cat Legends
Для початку гри переглянь інструкції:
/help
`

const (
	NewPlayerCallback       = "newPlayer"
	PlayerStatsCallback     = "playerStats"
	PlayerInventoryCallback = "playerInventory"
)

var newPlayerKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Створити персонажа", NewPlayerCallback),
	),
)

var existingPlayerKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Мій персонаж", PlayerStatsCallback),
		tgbotapi.NewInlineKeyboardButtonData("Інвентар", PlayerInventoryCallback),
	),
)

func Start(msg *tgbotapi.MessageConfig, update *tgbotapi.Update) {
	msg.Text = strings.Replace(startMessage, "%name%", update.Message.From.FirstName, 1)
	_, ok := game.GetPlayerById(update.Message.Chat.ID)
	if ok {
		msg.ReplyMarkup = existingPlayerKeyboard
	} else {
		msg.ReplyMarkup = newPlayerKeyboard
	}
}
