package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	controllerpb "happybirthday_bot/internal/controller"
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
	Process(bot *tgbotapi.BotAPI, update tgbotapi.Update) (err error)
}

func NewController() controller {
	return controllerpb.New()
}

func main() {
	secretDataStore := NewSecretDataStore()
	controller := NewController()

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
			err := controller.Process(bot, update)
			if err != nil {
				log.Printf("Authorized on account %s", bot.Self.UserName)
				tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла какая-то ошибка.\nМашины тоже ошибаются((")
			}
		}
	}
}
