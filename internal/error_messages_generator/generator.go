package error_messages_generator

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type errorGenerator struct{}

func New() *errorGenerator {
	return &errorGenerator{}
}

func (e *errorGenerator) GenerateErrorMessage(chatID int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, "какая-то ошибка")
	return msg
}
