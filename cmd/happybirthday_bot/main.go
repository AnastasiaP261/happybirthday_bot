package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	controllerpb "happybirthday_bot/internal/controller"
	"happybirthday_bot/internal/error_messages_generator"
	"happybirthday_bot/internal/secret_data_store"
	"log"
)

type secretDataGetter interface {
	BotToken() string
}

func NewSecretDataStore() secretDataGetter {
	return secret_data_store.New()
}

type controller interface {
	Process(update tgbotapi.Update) (msg tgbotapi.MessageConfig, err error)
}

func NewController() controller {
	return controllerpb.New()
}

type errorMsgGenerator interface {
	GenerateErrorMessage(chatID int64) tgbotapi.MessageConfig
}

func NewErrorMsgGenerator() errorMsgGenerator {
	return error_messages_generator.New()
}

func main() {
	secretDataStore := NewSecretDataStore()
	controller := NewController()
	errMsgGen := NewErrorMsgGenerator()

	bot, err := tgbotapi.NewBotAPI(secretDataStore.BotToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			msg, err := controller.Process(update)
			if err != nil {
				msg = errMsgGen.GenerateErrorMessage(update.Message.Chat.ID)
			}

			bot.Send(msg)
		}
	}
}
