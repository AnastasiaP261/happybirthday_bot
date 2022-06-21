package controller

import (
	b64 "encoding/base64"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	f "happybirthday_bot/internal/formatting"
	"log"
	"strings"
	"time"
)

func (p *process) Process(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	log.Printf("START - [%s] %s", update.Message.From.UserName, update.Message.Text)

	p.send(bot, "Дешифрую, жди...\n", time.Second*10, update.Message.Chat.ID, update.Message.MessageID)

	decodeBytes, err := b64.StdEncoding.DecodeString(update.Message.Text)
	if err != nil {
		log.Printf("ENCODE ERR - [%s] %s", update.Message.From.UserName, update.Message.Text)

		txt := p.errGen.GenerateErrorMessage(update.Message.Chat.ID)
		p.send(bot, txt, time.Second*1, update.Message.Chat.ID, -1)
		return nil
	}
	decodeText := strings.ToUpper(string(decodeBytes))

	answerText, err := p.process(decodeText)
	if err != nil {
		log.Printf("RULE ERR - [%s] %s", update.Message.From.UserName, decodeText)

		txt := p.errGen.GenerateErrorMessage(update.Message.Chat.ID)
		p.send(bot, txt, time.Second*1, update.Message.Chat.ID, -1)
		return err
	}

	p.send(bot, "Успешно!\nПолучено:\n\n"+f.ToBold(decodeText), time.Second*1, update.Message.Chat.ID, -1)
	p.send(bot, answerText, time.Second*1, update.Message.Chat.ID, -1)

	return nil
}

func (p *process) process(text string) (string, error) {
	switch text {
	default:
		return "", errors.New("invalid")
	}
}

// send msgID - необязательный параметр. При отправке ID сообщение отправится ответом на указанное. Если это не требуется, отправить -1
func (p *process) send(bot *tgbotapi.BotAPI, msgText string, sleepTime time.Duration, chatID int64, msgID int) {
	msg := tgbotapi.NewMessage(chatID, msgText)
	if msgID != -1 {
		msg.ReplyToMessageID = msgID
	}

	_, _ = bot.Send(msg)
	time.Sleep(sleepTime)
}
