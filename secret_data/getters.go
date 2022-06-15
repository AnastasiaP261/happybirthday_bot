package secret_data

type store struct{}

func New() *store {
	return &store{}
}

func (*store) BotToken() string {
	return botToken
}
