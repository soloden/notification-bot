package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/soloden/notificator-bot/pkg/logging"
)

type Config struct {
	BotToken string `yaml:"bot_token"`
	MongoDb  struct {
		DbUsername string `yaml:"username"`
		DbPassword string `yaml:"password"`
		DbHost     string `yaml:"host"`
		DbPort     string `yaml:"port"`
		DbName     string `yaml:"database"`
	}
}

var instance *Config
var once sync.Once

func LoadConfigs() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
