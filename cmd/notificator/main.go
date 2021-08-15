package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/soloden/notificator-bot/pkg/config"
	"github.com/soloden/notificator-bot/pkg/notificator"
)

func main() {
	cfg, err := config.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalln(err)
	}
	bot.Debug = true

	app := notificator.NewApp(bot)
	app.Start()
}
