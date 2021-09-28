package app

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/soloden/notificator-bot/internal/notification"
)

func (app *App) HandleStartCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Вы успешно авторизировались")

	app.bot.Send(msg)
}

func (app *App) HandleAddCommand(message *tgbotapi.Message, ctx context.Context) {
	ntfDto := notification.NotificationDto{}
	ntfDto.CreatedAt = message.Date
	ntfDto.OwnerId = message.Chat.ID
	ntfDto.Text = message.Text

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	id, err := app.service.Create(ctx, ntfDto)
	if err != nil {
		msg.Text = "Ошибка при создании уведомления"
		app.bot.Send(msg)
		return
	}

	nmsg := tgbotapi.NewMessage(message.Chat.ID, "Вы успешно создали напоминание, id - "+id)
	app.bot.Send(nmsg)
}

func (app *App) HandleListCommand(message *tgbotapi.Message, ctx context.Context) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Показать все напоминания")

	app.bot.Send(msg)
}

func (app *App) HandleUpdateCommand(message *tgbotapi.Message, ctx context.Context) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Обновить напоминалку")

	app.bot.Send(msg)
}

func (app *App) HandleDeleteCommand(message *tgbotapi.Message, ctx context.Context) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Удалить напоминалку")

	app.bot.Send(msg)
}

func (app *App) HandleUnknowCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда")
	app.bot.Send(msg)
}
