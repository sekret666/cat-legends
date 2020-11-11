package main

import (
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
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "start":
				msg.Text = "Hello"
			case "help":
				msg.Text = "help"
			case "echo":
				msg.Text = update.Message.CommandArguments()
			default:
				msg.Text = "I don't know that command. Please use /help to view all commands"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Error(err)
			}
		}

	}
}
