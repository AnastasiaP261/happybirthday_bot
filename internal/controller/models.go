package controller

import "happybirthday_bot/internal/secret_data_store"

type process struct {
	handlers []rule

	hints  []secret_data_store.Hint
	errGen errorMsgGenerator
}

func New(hints []secret_data_store.Hint, errGen errorMsgGenerator, rules ...rule) *process {
	return &process{
		errGen:   errGen,
		handlers: rules,
		hints:    hints,
	}
}

type errorMsgGenerator interface {
	GenerateErrorMessage() string
}

type rule interface {
	Hand() string
}

type Hint interface {
	GetWish() string
	GetCode() string
	GetHintText() string
}
