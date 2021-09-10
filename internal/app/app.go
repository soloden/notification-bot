package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/soloden/notificator-bot/internal/notification"
	"github.com/soloden/notificator-bot/pkg/logging"
)

type App struct {
	bot     *tgbotapi.BotAPI
	service *notification.Service
}

func NewApp(bot *tgbotapi.BotAPI, service *notification.Service) *App {
	return &App{bot: bot, service: service}
}

func (app *App) Start(logger *logging.Logger) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := app.bot.GetUpdatesChan(updateConfig)
	logger.Infoln("open chan for reciving message")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			app.ParseCommand(update.Message)
		}
	}
}
