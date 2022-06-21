package controller

type process struct {
	handlers []rule

	configs secretDataGetter
	errGen  errorMsgGenerator
	decoder decoder
}

func New(configs secretDataGetter, decoder decoder, errGen errorMsgGenerator, rules ...rule) *process {
	return &process{
		errGen:   errGen,
		decoder:  decoder,
		handlers: rules,
		configs:  configs,
	}
}

type secretDataGetter interface {
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
