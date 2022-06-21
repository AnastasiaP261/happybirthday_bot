package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	controllerpb "happybirthday_bot/internal/controller"
	"happybirthday_bot/internal/error_messages_generator"
	"happybirthday_bot/internal/rules"
	"happybirthday_bot/internal/secret_data_store"
	"log"
)

func main() {
	secretDataStore := secret_data_store.New()

	allRules := []controllerpb.Rule{
		rules.StartRule{},
	}

	controller := controllerpb.New(secretDataStore.GetHints(), error_messages_generator.New(), allRules)

	bot, err := tgbotapi.NewBotAPI(secretDataStore.GetBotToken())
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
