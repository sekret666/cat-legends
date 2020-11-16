package main

import (
	"CatLegends/events"
	"CatLegends/utils"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.SetFormatter(&utils.Formatter{})
	log.SetReportCaller(true)
}

func main() {
	utils.InitDB()
	defer utils.CloseDB()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Telegram Bot authorized: ", bot.Self.UserName)

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message != nil && update.Message.Chat.IsPrivate() {
			chatId := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatId, "")

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					events.Start(&msg, &update)
				case "help":
					events.Help(&msg)
				case "stats":
					events.Stats(&msg, chatId)
				case "inventory":
					events.Inventory(&msg, chatId, 0)
				default:
					msg.Text = events.UnknownCommandMessage
				}
			} else {
				msg.Text = events.UnknownMessage
			}

			if _, err := bot.Send(msg); err != nil {
				log.Error(err)
			}
		}

		if update.CallbackQuery != nil {
			chatId := update.CallbackQuery.Message.Chat.ID
			msgId := update.CallbackQuery.Message.MessageID
			cb := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			var msg tgbotapi.Chattable
			var ok bool

			qData := update.CallbackQuery.Data

			if qData == events.NewPlayerCallback {
				msg, ok = events.NewPlayer(&cb, chatId, msgId)
			} else if qData == events.PlayerStatsCallback {
				m := tgbotapi.NewMessage(chatId, "")
				events.Stats(&m, chatId)
				msg = m
				ok = true
			} else if qData == events.PlayerInventoryCallback {
				m := tgbotapi.NewMessage(chatId, "")
				events.Inventory(&m, chatId, 0)
				msg = m
				ok = true
			} else if strings.HasPrefix(qData, "page_") {
				page, err := strconv.ParseInt(qData[5:], 10, 32)
				if err != nil {
					log.Error(err)
				} else {
					msg, ok = events.UpdateInventory(msgId, chatId, int(page))
					if !ok {
						cb.Text = events.NoPlayerText
					}
				}
			} else if strings.HasPrefix(qData, "item_") {
				itemInd, err := strconv.ParseInt(qData[5:], 10, 32)
				if err != nil {
					log.Error(err)
				} else {
					fmt.Println(itemInd)
				}
			} else {
				cb.Text = events.UnknownCallback
			}

			if ok {
				if _, err := bot.Send(msg); err != nil {
					log.Error(err)
				}
			}

			_, err := bot.AnswerCallbackQuery(cb)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}
