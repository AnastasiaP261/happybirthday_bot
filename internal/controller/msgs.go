package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

// loading показывает сообщение с загрузочным "колесиком" в псевдографике в течение времени, которое установлено в таймере
func loading(bot *tgbotapi.BotAPI, chatID int64, timer *time.Timer) {
	pause := time.Millisecond * 200

	msg := tgbotapi.NewMessage(chatID, `\`)
	sended, _ := bot.Send(msg)

	alive := true
	for alive {
		select {
		case <-timer.C:
			alive = false
		default:
			time.Sleep(pause)
			bot.Send(tgbotapi.NewEditMessageText(chatID, sended.MessageID, `|`))
			time.Sleep(pause)
			bot.Send(tgbotapi.NewEditMessageText(chatID, sended.MessageID, `/`))
			time.Sleep(pause)
			bot.Send(tgbotapi.NewEditMessageText(chatID, sended.MessageID, `-`))
			time.Sleep(pause)
			bot.Send(tgbotapi.NewEditMessageText(chatID, sended.MessageID, `\`))
		}
	}

	bot.Send(tgbotapi.NewDeleteMessage(chatID, sended.MessageID))
}
