package main

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/soloden/notificator-bot/internal/app"
	"github.com/soloden/notificator-bot/internal/config"
	"github.com/soloden/notificator-bot/internal/notification"
	"github.com/soloden/notificator-bot/internal/notification/db"
	"github.com/soloden/notificator-bot/pkg/logging"
	mongo "github.com/soloden/notificator-bot/pkg/mongodb"
)

func main() {
	logger := logging.GetLogger()
	logger.Infoln("logger initialized")

	cfg := config.LoadConfigs()
	logger.Infoln("config initialized")

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		logger.Fatalf("tgbotapi create instance failed: %w", err)
	}
	bot.Debug = true
	logger.Infoln("telegram bot initialized")

	mongoClient, err := mongo.NewClient(context.Background(), cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	if err != nil {
		log.Fatal(err)
	}
	logger.Infoln("mongo client initialized")

	ntfStorage := db.NewStorage(mongoClient, cfg.DbName, logger)
	logger.Infoln("notification storage initialized")

	ntfService := notification.NewService(ntfStorage, *logger)
	logger.Infoln("notification service initialized")

	mainApp := app.NewApp(bot, ntfService)
	mainApp.Start(logger)
}
