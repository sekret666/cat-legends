package main

import (
	"CatLegends/utils"
	"context"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
	"time"
)

func init() {
	log.SetFormatter(&utils.Formatter{})
	log.SetReportCaller(true)
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	//collection := client.Database("cat_legends").Collection("entities")

	//p := game.NewPlayer()
	//p.Level = game.Level{
	//	Level:       5,
	//	XP:  2,
	//	LevelUpXP: 7,
	//}
	//p.Health = game.Health{
	//	Health:    20,
	//	MaxHealth: 24,
	//}
	//p.Mana = game.Mana{
	//	Mana:    8,
	//	MaxMana: 12,
	//}
	//
	//insertResult, err := collection.InsertOne(context.TODO(), p)
	//if err != nil {
	//	log.Error(err)
	//}
	//fmt.Println("Inserted with ID:", insertResult.InsertedID)

	//filter := bson.D{}
	//var newP game.Player
	//err = collection.FindOne(context.TODO(), filter).Decode(&newP)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Found: ", &newP)

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

		if update.Message.Chat.IsPrivate() {
			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				switch update.Message.Command() {
				case "start":
					msg.Text = strings.Replace(startMessage, "%name%", update.Message.From.FirstName, 1)
				case "help":
					msg.Text = helpMessage
				default:
					msg.Text = unknownCommandMessage
				}

				if _, err := bot.Send(msg); err != nil {
					log.Error(err)
				}
			}
		}
	}
}
