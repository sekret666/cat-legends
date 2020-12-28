package events

import (
	"CatLegends/game"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const playerStats = `
<code>🎖 Рівень: %lvl%
✨ Досвід: %XP%/%lvlUpXP%</code>

<code>❤️ Здоров'я: %HP%/%maxHP%
🔵 Мана:     %MP%/%maxMP%</code>

<code>✊ Витривалість: %end%
🏃🏻 Спритність:   %agi%
🧠 Інтелект:     %int%
💪 Сила:         %str%
🍀 Удача:        %luc%</code>
`

func Stats(msg *tgbotapi.MessageConfig, chatId int64) {
	p, ok := game.GetPlayerById(chatId)
	if ok {
		msgText := p.Level.ReplaceInString(playerStats)
		msgText = p.Health.ReplaceInString(msgText)
		msgText = p.Mana.ReplaceInString(msgText)
		msgText = p.Attributes.ReplaceInString(msgText)

		msg.Text = msgText
		msg.ParseMode = tgbotapi.ModeHTML
	} else {
		msg.Text = NoPlayerText
	}
}
