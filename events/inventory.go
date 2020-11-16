package events

import (
	"CatLegends/game"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const playerInventory = `
<code>💰 Гроші: %money%</code>

Ваші речі:
`

func Inventory(msg *tgbotapi.MessageConfig, chatId int64, page int) {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		msgText := p.Inventory.Money.ReplaceInString(playerInventory)

		msg.Text = msgText
		msg.ParseMode = tgbotapi.ModeHTML

		msg.ReplyMarkup = p.Inventory.GetInlineKeyboard(page)
	} else {
		msg.Text = NoPlayerText
	}
}

func UpdateInventory(msgId int, chatId int64, page int) (tgbotapi.Chattable, bool) {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		msgEdit := tgbotapi.NewEditMessageReplyMarkup(chatId, msgId, p.Inventory.GetInlineKeyboard(page))
		return msgEdit, true
	} else {
		return nil, false
	}
}