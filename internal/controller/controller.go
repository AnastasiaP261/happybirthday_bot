package controller

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	f "happybirthday_bot/internal/formatting"
	"happybirthday_bot/internal/secret_data_store"
	"log"
	"time"
)

func (p *process) Process(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	log.Printf("START - [%s] %s", update.Message.From.UserName, update.Message.Text)

	if len(update.Message.Entities) != 0 && update.Message.Entities[0].Type == "bot_command" {
		if update.Message.Text == "/start" {
			p.rules[0].RuleProcessing(bot, update.Message.Chat.ID)
			return nil
		}
	}

	p.send(bot, "Дешифрую, жди\\.\\.\\.\n", time.Second*10, update.Message.Chat.ID, update.Message.MessageID)

	hint, err := p.decode(update.Message.Text)
	if err != nil {
		log.Printf("DECODE ERR - [%s] %s", update.Message.From.UserName, update.Message.Text)

		txt := p.errGen.GenerateErrorMessage()
		p.send(bot, txt, time.Second*1, update.Message.Chat.ID, -1)
		return nil
	}

	p.send(bot, "Успешно\\!\nТебе пожелали:\n\n"+f.ToBold(hint.Wish), time.Second*3, update.Message.Chat.ID, -1)
	p.send(bot, hint.GetHintText(), time.Second*10, update.Message.Chat.ID, -1)

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
