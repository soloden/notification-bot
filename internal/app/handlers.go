package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (app *App) HandleStartCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Вы успешно авторизировались")

	app.bot.Send(msg)
}

func (app *App) HandleAddCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Добавить новое напомнинание")

	app.bot.Send(msg)
}

func (app *App) HandleListCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Показать все напоминания")

	app.bot.Send(msg)
}

func (app *App) HandleUpdateCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Обновить напоминалку")

	app.bot.Send(msg)
}

func (app *App) HandleDeleteCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Удалить напоминалку")

	app.bot.Send(msg)
}

func (app *App) HandleUnknowCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда")
	app.bot.Send(msg)
}
