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
				case "battle":
					events.Battle(&msg, chatId)
				default:
					msg.Text = events.UnknownCommandMessage
				}
			} else if update.Message.Dice != nil {
				switch update.Message.Dice.Emoji {
				case "🎲":
					events.EscapeStatus(&msg, &update)
				default:
					msg.Text = events.UnknownMessage
				}
			} else {
				msg.Text = events.UnknownMessage
			}

			if _, err := bot.Send(msg); err != nil {
				log.Error(err)
			}

			continue
		}

		if update.CallbackQuery != nil && update.CallbackQuery.Message != nil && update.CallbackQuery.Message.Chat.IsPrivate() {
			chatId := update.CallbackQuery.Message.Chat.ID
			msgId := update.CallbackQuery.Message.MessageID

			cb := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			msg := tgbotapi.NewMessage(chatId, "")

			queryData := strings.Split(update.CallbackQuery.Data, "_")

			switch queryData[0] {
			case events.NewPlayerCallback:
				events.NewPlayer(&cb, chatId, msgId, bot)
			case events.PlayerStatsCallback:
				events.Stats(&msg, chatId)
				cb.Text = "Статистика"
			case events.PlayerInventoryCallback:
				events.Inventory(&msg, chatId, 0)
				cb.Text = "Інвентар"
			case events.NewBattleCallback:
				cb.Text = "Бій"
			case events.NewEscapeCallback:
				events.Escape(&msg)
				cb.Text = "Втеча"
			case "page":
				page, err := strconv.ParseInt(queryData[1], 10, 32)
				if err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				} else {
					ok := events.UpdateInventory(msgId, chatId, int(page), bot)
					if !ok {
						cb.Text = events.NoPlayerText
					}
				}
			case "item":
				itemInd, err := strconv.ParseInt(queryData[1], 10, 32)
				if err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				} else {
					ok := events.ShowItem(msgId, chatId, int(itemInd), bot)
					if !ok {
						cb.Text = events.ErrorText
					}
				}
			default:
				cb.Text = events.UnknownCallback
			}

			if msg.Text != "" {
				if _, err := bot.Send(msg); err != nil {
					log.Error(err)
					cb.Text = events.ErrorText
				}
			}

			if _, err := bot.Send(cb); err != nil {
				log.Error(err)
			}

			continue
		}
	}
}
