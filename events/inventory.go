package events

import (
	"CatLegends/game"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const playerInventory = `
<code>💰 Гроші: %money%</code>

Ваші речі:
`

func Inventory(msg *tgbotapi.MessageConfig, chatId int64) {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		msgText := p.Inventory.Money.ReplaceInString(playerInventory)

		msg.Text = msgText
		msg.ParseMode = tgbotapi.ModeHTML

		msg.ReplyMarkup = p.Inventory.GetInlineKeyboard(0)
	} else {
		msg.Text = NoPlayerText
	}
}
