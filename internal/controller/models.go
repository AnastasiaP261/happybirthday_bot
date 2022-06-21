package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"happybirthday_bot/internal/secret_data_store"
)

type process struct {
	rules []Rule

	hints  []secret_data_store.Hint
	errGen errorMsgGenerator
}

func New(hints []secret_data_store.Hint, errGen errorMsgGenerator, rules []Rule) *process {
	return &process{
		errGen: errGen,
		rules:  rules,
		hints:  hints,
	}
}

type errorMsgGenerator interface {
	GenerateErrorMessage() string
}

type Rule interface {
	RuleProcessing(bot *tgbotapi.BotAPI, chatID int64)
}

type Hint interface {
	GetWish() string
	GetCode() string
	GetHintText() string
	GetNum() int64
}
