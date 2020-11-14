package events

import (
	"CatLegends/game"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func Stats(msg *tgbotapi.MessageConfig, chatId int64) {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		msg.Text = p.String()
	} else {
		msg.Text = "No player found!"
	}
}
