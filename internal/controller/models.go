package controller

import (
	"happybirthday_bot/internal/error_messages_generator"
)

type process struct {
	handlers []rule

	errGen errorMsgGenerator
}

func New(hadlers ...rule) *process {
	return &process{
		errGen:   NewErrorMsgGenerator(),
		handlers: hadlers,
	}
}

type errorMsgGenerator interface {
	GenerateErrorMessage(chatID int64) string
}

func NewErrorMsgGenerator() errorMsgGenerator {
	return error_messages_generator.New()
}

type rule interface {
	Hand() string
}
