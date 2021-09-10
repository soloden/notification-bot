package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/soloden/notificator-bot/pkg/logging"
)

type Config struct {
	BotToken   string `env:"BOT_TOKEN"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_DATABASE"`
}

var instance *Config
var once sync.Once

func LoadConfigs() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
