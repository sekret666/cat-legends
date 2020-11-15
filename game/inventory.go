package game

import "github.com/go-telegram-bot-api/telegram-bot-api"

const itemsPerRow = 4
const rowsRerPage = 2

type Inventory struct {
	Money Money  `bson:"money"`
	Items []Item `bson:"items"`
}

func (inv *Inventory) GetInlineKeyboard(page int) tgbotapi.InlineKeyboardMarkup {
	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(inv.Items[0].Emoji, "item1"),
			tgbotapi.NewInlineKeyboardButtonData(inv.Items[1].Emoji, "item2")),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<", "back"),
			tgbotapi.NewInlineKeyboardButtonData(">", "next")))

	//for i := 0; i < rowsRerPage; i++ {
	//	for j := 0; j < itemsPerRow; j++ {
	//		ind := itemsPerRow * i + j;
	//		if ind < len(inv.Items) {
	//
	//		}
	//	}
	//}

	return kb
}
