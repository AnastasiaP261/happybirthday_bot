package secret_data_store

type store struct{}

func New() *store {
	return &store{}
}

func (*store) BotToken() string {
	return botToken
}
