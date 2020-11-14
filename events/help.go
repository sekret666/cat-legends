package events

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpMessage = `
Як почати грати? Ось інструкції...
`

func Help(msg *tgbotapi.MessageConfig) {
	msg.Text = helpMessage
}
