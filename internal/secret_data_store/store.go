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
	Wishes []string `yaml:"wishes"`
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
