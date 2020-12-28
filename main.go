package main

import (
	"CatLegends/events"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	log.Info("Startup")
	log.Info("------------------")

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

	updates := bot.GetUpdatesChan(u)

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

			qData := update.CallbackQuery.Data

			if qData == events.NewPlayerCallback {
				events.NewPlayer(&cb, chatId, msgId, bot)
			} else if qData == events.PlayerStatsCallback {
				msg := tgbotapi.NewMessage(chatId, "")
				events.Stats(&msg, chatId)
				if _, err := bot.Send(msg); err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				}
			} else if qData == events.PlayerInventoryCallback {
				msg := tgbotapi.NewMessage(chatId, "")
				events.Inventory(&msg, chatId, 0)
				if _, err := bot.Send(msg); err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				}
			} else if strings.HasPrefix(qData, "page_") {
				page, err := strconv.ParseInt(qData[5:], 10, 32)
				if err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				} else {
					ok := events.UpdateInventory(msgId, chatId, int(page), bot)
					if !ok {
						cb.Text = events.NoPlayerText
					}
				}
			} else if strings.HasPrefix(qData, "item_") {
				itemInd, err := strconv.ParseInt(qData[5:], 10, 32)
				if err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				} else {
					ok := events.ShowItem(msgId, chatId, int(itemInd), bot)
					if !ok {
						cb.Text = events.ErrorText
					}
				}
			} else {
				cb.Text = events.UnknownCallback
			}

			_, err := bot.Send(cb)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}
