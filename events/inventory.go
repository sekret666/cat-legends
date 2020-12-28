package events

import (
	"CatLegends/game"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const playerInventory = `
<code>💰 Гроші: %money%</code>

Ваші речі:
`

const itemInfo = `
%itemEmoji% <code>%itemName%</code> (%itemQuantity%)
<i>%itemRarity%</i>

<code>%itemDescription%</code>

%itemTypeInfo%

Ціна: %itemPrice%
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

func UpdateInventory(msgId int, chatId int64, page int, bot *tgbotapi.BotAPI) bool {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		msgEdit := tgbotapi.NewEditMessageText(chatId, msgId, "")

		msgText := p.Inventory.Money.ReplaceInString(playerInventory)
		msgEdit.Text = msgText
		msgEdit.ParseMode = tgbotapi.ModeHTML

		inventoryKeyboard := p.Inventory.GetInlineKeyboard(page)
		msgEdit.ReplyMarkup = &inventoryKeyboard

		if _, err := bot.Send(msgEdit); err != nil {
			log.Error(err)
			return false
		}
		return true
	} else {
		return false
	}
}

func ShowItem(msgId int, chatId int64, itemInd int, bot *tgbotapi.BotAPI) bool {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		if itemInd >= len(p.Inventory.Items) {
			return false
		}
		item := p.Inventory.Items[itemInd]

		msgEdit := tgbotapi.NewEditMessageText(chatId, msgId, item.ReplaceInString(itemInfo))
		msgEdit.ParseMode = tgbotapi.ModeHTML

		page := game.GetInventoryPageFromIndex(itemInd)
		backKeyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Назад", "page_"+strconv.Itoa(page))))
		msgEdit.ReplyMarkup = &backKeyboard

		if _, err := bot.Send(msgEdit); err != nil {
			log.Error(err)
			return false
		}
		return true
	} else {
		return false
	}
}
