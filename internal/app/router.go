package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (app *App) ParseCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		app.HandleStartCommand(message)
	case "add":
		app.HandleAddCommand(message)
	case "update":
		app.HandleAddCommand(message)
	case "delete":
		app.HandleAddCommand(message)
	case "list":
		app.HandleAddCommand(message)
	default:
		app.HandleUnknowCommand(message)
	}

}
