package events

import (
	"CatLegends/game"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
)

const (
	escapeText       = `Щоб втекти киньте <code>🎲</code>, щоб утекти потрібно 4 і більше`
	escapeStatusText = `Випало %dice%, ви %status%`
)

func Escape(msg *tgbotapi.MessageConfig) {
	msg.Text = escapeText
	msg.ParseMode = tgbotapi.ModeHTML
}

func EscapeStatus(msg *tgbotapi.MessageConfig, update *tgbotapi.Update) {
	chatId := update.Message.Chat.ID

	e, ok := game.GetEnemyById(chatId)
	if ok {
		if e.EscapeStatus == game.CanEscape {
			dice := update.Message.Dice.Value

			msgText := strings.ReplaceAll(escapeStatusText, "%dice%", strconv.Itoa(dice))
			if dice >= 4 {
				msgText = strings.ReplaceAll(msgText, "%status%", "втекли")
			} else {
				msgText = strings.ReplaceAll(msgText, "%status%", "не змогли втекли")
				e.EscapeStatus = game.CantEscape

				db := utils.GetDB()
				eUpdate := bson.M{
					"$set": bson.M{
						"escapeStatus": int(e.EscapeStatus),
					},
				}
				_, err := db.Enemies.UpdateOne(db.Ctx, bson.M{"chatId": chatId}, eUpdate)
				if err != nil {
					log.Error(err)
				}
			}

			msg.Text = msgText
		} else if e.EscapeStatus == game.CantEscape {
			msg.Text = "Ви більше не можете втекти"
		}
	} else {
		msg.Text = UnknownMessage
	}
}
