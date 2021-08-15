package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

func LoadConfigs() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg.BotToken = os.Getenv("BOT_TOKEN")

	return &cfg, nil
}
