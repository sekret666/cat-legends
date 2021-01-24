package events

import (
	"CatLegends/game"
	"CatLegends/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

const startBattle = `
<i>%name%</i> Lvl %lvl%
<code>❤ ️%HP%/%maxHP%</code>

<i>%description%</i>
`

const (
	NewBattleCallback = "newBattle"
	NewEscapeCallback = "newEscape"
)

var newEnemyKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Битися", NewBattleCallback),
		tgbotapi.NewInlineKeyboardButtonData("Втекти", NewEscapeCallback),
	),
)

func Battle(msg *tgbotapi.MessageConfig, chatId int64) {
	db := utils.GetDB()

	var e *game.Enemy
	var ok bool

	e, ok = game.GetEnemyById(chatId)
	if !ok {
		e = game.GenerateEnemy(chatId, 1)
		_, err := db.Enemies.InsertOne(db.Ctx, e)
		if err != nil {
			log.Error(err)
			return
		}
	}

	msgText := e.ReplaceInString(startBattle)
	msgText = e.Level.ReplaceInString(msgText)
	msgText = e.Health.ReplaceInString(msgText)

	msg.Text = msgText
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = newEnemyKeyboard
}
