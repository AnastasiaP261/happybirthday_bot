package rules

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"time"
)

type StartRule struct{}

func (r StartRule) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	data, _ := ioutil.ReadFile("images/pila.jpg")
	b := tgbotapi.FileBytes{Name: "pila.jpg", Bytes: data}

	msg := tgbotapi.NewPhoto(chatID, b)
	msg.Caption = hello
	msg.ParseMode = "MarkdownV2"

	_, _ = bot.Send(msg)
	time.Sleep(time.Second * 25)
}
