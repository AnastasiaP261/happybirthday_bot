package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	controllerpb "happybirthday_bot/internal/controller"
	"happybirthday_bot/internal/error_messages_generator"
	"happybirthday_bot/internal/hints"
	"happybirthday_bot/internal/secret_data_store"
	"log"
	"sync"
)

func main() {
	secretDataStore := secret_data_store.New()

	allHints := []controllerpb.Rule{
		rules.StartHint{},
		rules.Hint1{},
		rules.Hint2{},
		rules.Hint3{},
		rules.Hint4{},
		rules.Hint5{},
		rules.Hint6{},
		rules.Hint7{},
		rules.Hint8{},
		rules.Hint9{},
		rules.Hint10{},
	}

	controller := controllerpb.New(secretDataStore.GetHints(), error_messages_generator.New(), allHints)

	bot, err := tgbotapi.NewBotAPI(secretDataStore.GetBotToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var wg sync.WaitGroup
	goroutines := make(chan struct{}, 1) // ограничение кол-ва потоков по необходимости
	for update := range updates {
		update := update
		goroutines <- struct{}{}
		wg.Add(1)
		go func(goroutines <-chan struct{}) {
			if update.Message != nil { // If we got a message
				err := controller.Process(bot, update)
				if err != nil {
					log.Printf("Authorized on account %s", bot.Self.UserName)
					tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла какая-то ошибка.\nМашины тоже ошибаются((")
				}
				<-goroutines
				wg.Done()
			}
		}(goroutines)
	}
}
