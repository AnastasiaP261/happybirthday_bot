package error_messages_generator

type ErrorGenerator struct{}

func New() *ErrorGenerator {
	return &ErrorGenerator{}
}

func (e *ErrorGenerator) GenerateErrorMessage() string {
	return "какая-то ошибка"
}
