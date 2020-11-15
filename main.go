package main

import (
	"CatLegends/events"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"os"
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
					events.Inventory(&msg, chatId)
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

			switch update.CallbackQuery.Data {
			case events.NewPlayerCallback:
				msg, ok = events.NewPlayer(&cb, chatId, msgId)
			case events.ExistingPlayerCallback:
				m := tgbotapi.NewMessage(chatId, "")
				events.Stats(&m, chatId)
				msg = m
				ok = true
			default:
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
