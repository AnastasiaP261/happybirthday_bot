package secret_data_store

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type store struct {
	Bot struct {
		BotToken string `yaml:"botToken"`
	} `yaml:"bot"`

	Hints []Hint `yaml:"hints"`

	Wishes []string `yaml:"wishes"`
}

type Hint struct {
	Wish     string `yaml:"wish"`
	Code     string `yaml:"code"`
	HintText string `yaml:"hintText"`
	Num      int64  `yaml:"num"`
}

func (h Hint) GetWish() string {
	return h.Wish
}

func (h Hint) GetCode() string {
	return h.Code
}

func (h Hint) GetHintText() string {
	return h.HintText
}

func (h Hint) GetNum() int64 {
	return h.Num
}

func New() *store {
	secrDS := &store{}
	secrDS.getConf()

	return secrDS
}

func (s *store) getConf() *store {
	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}

func (s *store) GetBotToken() string {
	return s.Bot.BotToken
}

func (s *store) GetWishes() []string {
	return s.Wishes
}

func (s *store) GetHints() []Hint {
	return s.Hints
}
