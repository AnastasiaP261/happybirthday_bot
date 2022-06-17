package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type process struct {
}

func New() *process {
	return &process{}
}

func (p *process) Process(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	return msg, nil
}
