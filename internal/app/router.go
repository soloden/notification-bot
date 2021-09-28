package app

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (app *App) ParseCommand(message *tgbotapi.Message) {

	ctx := context.Background()

	switch message.Command() {
	case "start":
		app.HandleStartCommand(message)
	case "add":
		app.HandleAddCommand(message, ctx)
	case "update":
		app.HandleListCommand(message, ctx)
	case "delete":
		app.HandleUpdateCommand(message, ctx)
	case "list":
		app.HandleDeleteCommand(message, ctx)
	default:
		app.HandleUnknowCommand(message)
	}

}
