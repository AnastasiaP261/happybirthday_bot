package error_messages_generator

import "math/rand"

type ErrorGenerator struct{}

func New() *ErrorGenerator {
	return &ErrorGenerator{}
}

var errTexts = []string{"Какая-то ошибка(", "Проверь еще раз", "Ошибочка", "Ты пытаешься меня обмануть?", "Попробуй еще",
	"Не удалось", "Неа", "Снова ошибка", "Опять ошибка", "И еще разок попробуй", "Перепроверь", "Что-то не то", "Неть.",
	"Не удалось", "Давай еще раз", "Ошибся ты, а об ошибке пишу я", "Ошиииибка", "Больше ошибок богу ошибок!!!",
	"О Ш И Б К А", "Подумай еще", "Не то", "Не так", "Еще ошибка", "Ошибка", "Все когда-то ошибаются",
}

func (e *ErrorGenerator) GenerateErrorMessage() string {
	return errTexts[rand.Int()%len(errTexts)]
}
