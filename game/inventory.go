package game

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const itemsPerRow = 4
const rowsRerPage = 2

type Inventory struct {
	Money Money  `bson:"money"`
	Items []Item `bson:"items"`
}

func (inv *Inventory) GetInlineKeyboard(page int) tgbotapi.InlineKeyboardMarkup {
	kb := tgbotapi.NewInlineKeyboardMarkup()

	itemCount := len(inv.Items)
	offset := page * rowsRerPage * itemsPerRow
	if itemCount > offset {
		for i := 0; i < rowsRerPage; i++ {
			if itemCount-offset-i*itemsPerRow > 0 {
				kb.InlineKeyboard = append(kb.InlineKeyboard, tgbotapi.NewInlineKeyboardRow())
				for j := 0; j < itemsPerRow; j++ {
					ind := itemsPerRow*i + j + offset
					if ind < len(inv.Items) {
						item := inv.Items[ind]
						btnText := item.Emoji
						if item.Quantity > 1 {
							btnText += fmt.Sprintf(" (%d)", item.Quantity)
						}
						kb.InlineKeyboard[i] = append(kb.InlineKeyboard[i], tgbotapi.NewInlineKeyboardButtonData(btnText, "item_"+strconv.Itoa(ind)))
					} else {
						break
					}
				}
			} else {
				break
			}
		}
	}

	nextBackRow := tgbotapi.NewInlineKeyboardRow()
	if page > 0 {
		nextBackRow = append(nextBackRow, tgbotapi.NewInlineKeyboardButtonData("Назад", "page_"+strconv.Itoa(page-1)))
	}
	if offset+rowsRerPage*itemsPerRow < itemCount {
		nextBackRow = append(nextBackRow, tgbotapi.NewInlineKeyboardButtonData("Вперед", "page_"+strconv.Itoa(page+1)))
	}

	if len(nextBackRow) > 0 {
		kb.InlineKeyboard = append(kb.InlineKeyboard, nextBackRow)
	}

	return kb
}

func GetInventoryPageFromIndex(ind int) int {
	return ind / (itemsPerRow * rowsRerPage)
}