package notificator

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type App struct {
	bot *tgbotapi.BotAPI
}

func NewApp(bot *tgbotapi.BotAPI) *App {
	return &App{bot: bot}
}

func (app *App) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := app.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		app.bot.Send(msg)
	}
}
