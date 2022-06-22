package controller

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	f "happybirthday_bot/internal/formatting"
	"happybirthday_bot/internal/secret_data_store"
	"log"
	"regexp"
	"time"
)

func (p *process) Process(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	log.Printf("START - [%s] %s", update.Message.From.UserName, update.Message.Text)

	// сначала пытаемся распознать команду
	if len(update.Message.Entities) != 0 && update.Message.Entities[0].Type == "bot_command" {
		if update.Message.Text == "/start" {
			p.rules[0].RuleProcessing(bot, update.Message.Chat.ID)
			return nil
		}
	}

	// проверяем похоже ли сообщение на код. Если нет - говорим об ошибке
	if ok, err := regexp.Match(`^[\w]{10}$`, []byte(update.Message.Text)); err != nil || !ok {
		log.Printf("INPUT ERR - [%s] %s", update.Message.From.UserName, update.Message.Text)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, p.errGen.GenerateErrorMessage()))
		return nil
	}

	// говорим о начале дешифрации
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Дешифрую, жди...\n")
	msg.ReplyToMessageID = update.Message.MessageID
	_, _ = bot.Send(msg)

	loading(bot, update.Message.Chat.ID, time.NewTimer(time.Second*5))

	// смотрим есть ли такой код в сценарии и вытаскиваем пожелание либо отправляем ошибку
	hint, err := p.decode(update.Message.Text)
	if err != nil {
		log.Printf("DECODE ERR - [%s] %s", update.Message.From.UserName, update.Message.Text)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, p.errGen.GenerateErrorMessage()))
		return nil
	}

	// отправляем сообщение об успехе и подсказку
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Успешно!\nТебе пожелали:\n\n"+f.ToBold(hint.Wish))
	msg.ParseMode = "MarkdownV2"
	bot.Send(msg)
	time.Sleep(time.Second * 1)

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, hint.GetHintText()))
	time.Sleep(time.Second * 1)

	return nil
}

func (p *process) decode(msgText string) (secret_data_store.Hint, error) {
	for _, hint := range p.hints {
		if hint.GetCode() == f.Normalize(msgText) {
			return hint, nil
		}
	}
	return secret_data_store.Hint{}, errors.New("no match")
}

// send При отправке msgID сообщение отправится ответом на указанное. Если это не требуется, отправить -1
func (p *process) send(bot *tgbotapi.BotAPI, msgText string, sleepTime time.Duration, chatID int64, msgID int) {
	msg := tgbotapi.NewMessage(chatID, msgText)
	if msgID != -1 {
		msg.ReplyToMessageID = msgID
	}
	msg.ParseMode = "MarkdownV2"

	_, _ = bot.Send(msg)
	time.Sleep(sleepTime)
}
