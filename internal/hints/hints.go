package rules

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"time"
)

type StartHint struct{}

func (r StartHint) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	data, _ := ioutil.ReadFile("images/pila.jpg")
	b := tgbotapi.FileBytes{Name: "pila.jpg", Bytes: data}

	msg := tgbotapi.NewPhoto(chatID, b)
	msg.Caption = hello
	msg.ParseMode = "MarkdownV2"

	_, _ = bot.Send(msg)
}

type Hint1 struct{}

func (r Hint1) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		msg := tgbotapi.NewMessage(chatID, hint1_1)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		data, _ := ioutil.ReadFile("images/putin.jpg")
		b := tgbotapi.FileBytes{Name: "putin.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		msg.Caption = hint1_2
		msg.ParseMode = "MarkdownV2"

		bot.Send(msg)
	}
}

type Hint2 struct{}

func (r Hint2) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	data, _ := ioutil.ReadFile("images/tort.jpg")
	b := tgbotapi.FileBytes{Name: "tort.jpg", Bytes: data}

	msg := tgbotapi.NewPhoto(chatID, b)
	msg.Caption = hint2
	msg.ParseMode = "MarkdownV2"

	bot.Send(msg)
}

type Hint3 struct{}

func (r Hint3) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		data, _ := ioutil.ReadFile("images/ivan.jpg")
		b := tgbotapi.FileBytes{Name: "ivan.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		msg.Caption = hint3_1
		msg.ParseMode = "MarkdownV2"

		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		msg := tgbotapi.NewMessage(chatID, hint3_2)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
}

type Hint4 struct{}

func (r Hint4) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		txts := []string{hint4_1, hint4_2, hint4_3, hint4_4, hint4_5, hint4_6}

		for _, t := range txts {
			msg := tgbotapi.NewMessage(chatID, t)
			msg.ParseMode = "MarkdownV2"
			bot.Send(msg)
			time.Sleep(time.Second * 1)
		}
	}
	time.Sleep(time.Second * 20)
	{
		msg := tgbotapi.NewMessage(chatID, hint4_7)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
		time.Sleep(time.Millisecond * 50)
	}
	{
		data, _ := ioutil.ReadFile("images/avocado.jpg")
		b := tgbotapi.FileBytes{Name: "avocado.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		bot.Send(msg)
	}
}

type Hint5 struct{}

func (r Hint5) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		txts := []string{hint5_1, hint5_2, hint5_3, hint5_4, hint5_5, hint5_6, hint5_7}

		for _, t := range txts {
			msg := tgbotapi.NewMessage(chatID, t)
			msg.ParseMode = "MarkdownV2"
			bot.Send(msg)
			time.Sleep(time.Millisecond * 150)
		}
	}
	time.Sleep(time.Second * 9)
	{
		msg := tgbotapi.NewMessage(chatID, hint5_8)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		data, _ := ioutil.ReadFile("images/ogurec.jpg")
		b := tgbotapi.FileBytes{Name: "ogurec.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		msg := tgbotapi.NewMessage(chatID, hint5_9)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
}

type Hint6 struct{}

func (r Hint6) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		msg := tgbotapi.NewMessage(chatID, hint6_1)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 3)
	{
		data, _ := ioutil.ReadFile("images/sestra.jpg")
		b := tgbotapi.FileBytes{Name: "sestra.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		msg.Caption = hint6_2
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
}

type Hint7 struct{}

func (r Hint7) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	data, _ := ioutil.ReadFile("images/vopros.jpg")
	b := tgbotapi.FileBytes{Name: "vopros.jpg", Bytes: data}

	msg := tgbotapi.NewPhoto(chatID, b)
	msg.Caption = hint7
	msg.ParseMode = "MarkdownV2"
	bot.Send(msg)
}

type Hint8 struct{}

func (r Hint8) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		data, _ := ioutil.ReadFile("images/gramota.jpg")
		b := tgbotapi.FileBytes{Name: "gramota.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		msg.Caption = hint8_1
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		msg := tgbotapi.NewMessage(chatID, hint8_2)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
}

type Hint9 struct{}

func (r Hint9) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		msg := tgbotapi.NewMessage(chatID, hint9_1)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 3)
	{
		msg := tgbotapi.NewMessage(chatID, hint9_2)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
}

type Hint10 struct{}

func (r Hint10) RuleProcessing(bot *tgbotapi.BotAPI, chatID int64) {
	{
		data, _ := ioutil.ReadFile("images/kaif.jpg")
		b := tgbotapi.FileBytes{Name: "kaif.jpg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		txts := []string{hint10_1, hint10_2, hint10_3, hint10_4}

		for _, t := range txts {
			msg := tgbotapi.NewMessage(chatID, t)
			msg.ParseMode = "MarkdownV2"
			bot.Send(msg)
			time.Sleep(time.Second * 1)
		}
	}
	time.Sleep(time.Second * 20)
	{
		txts := []string{hint10_5, hint10_6, hint10_7}

		for _, t := range txts {
			msg := tgbotapi.NewMessage(chatID, t)
			msg.ParseMode = "MarkdownV2"
			bot.Send(msg)
			time.Sleep(time.Second * 1)
		}
	}
	time.Sleep(time.Second * 1)
	{
		data, _ := ioutil.ReadFile("images/brat.jpeg")
		b := tgbotapi.FileBytes{Name: "brat.jpeg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		bot.Send(msg)
	}

	r.finish(bot, chatID)
}

func (r Hint10) finish(bot *tgbotapi.BotAPI, chatID int64) {
	time.Sleep(time.Minute * 10)

	{
		msg := tgbotapi.NewMessage(chatID, finish1)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 5)
	{
		msg := tgbotapi.NewMessage(chatID, finish2)
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	}
	time.Sleep(time.Second * 1)
	{
		msg := tgbotapi.NewMessage(chatID, finish3)
		msg.ParseMode = "MarkdownV2"
		sended, _ := bot.Send(msg)

		time.Sleep(time.Second * 1)

		bot.Send(tgbotapi.NewDeleteMessage(chatID, sended.MessageID))
	}
	{
		txts := []string{hint10_1, hint10_2, hint10_3, hint10_4}

		for _, t := range txts {
			msg := tgbotapi.NewMessage(chatID, t)
			msg.ParseMode = "MarkdownV2"
			bot.Send(msg)
			time.Sleep(time.Second * 1)
		}
	}
	time.Sleep(time.Second * 20)
	{
		txts := []string{hint10_5, hint10_6, hint10_7}

		for _, t := range txts {
			msg := tgbotapi.NewMessage(chatID, t)
			msg.ParseMode = "MarkdownV2"
			bot.Send(msg)
			time.Sleep(time.Second * 1)
		}
	}
	time.Sleep(time.Second * 1)
	{
		data, _ := ioutil.ReadFile("images/brat.jpeg")
		b := tgbotapi.FileBytes{Name: "brat.jpeg", Bytes: data}

		msg := tgbotapi.NewPhoto(chatID, b)
		bot.Send(msg)
	}
	time.Sleep(time.Minute * 10)
}
