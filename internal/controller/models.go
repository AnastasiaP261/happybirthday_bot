package controller

type process struct {
	handlers []rule

	errGen  errorMsgGenerator
	decoder decoder
}

func New(decoder decoder, errGen errorMsgGenerator, rules ...rule) *process {
	return &process{
		errGen:   errGen,
		decoder:  decoder,
		handlers: rules,
	}
}

type errorMsgGenerator interface {
	GenerateErrorMessage(chatID int64) string
}

type decoder interface {
	Decode(str string) (string, error)
}

type rule interface {
	Hand() string
}
