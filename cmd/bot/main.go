package main

import (
	"context"
	"os"

	"github.com/moguchev/telegram-bot/internal/app/bot"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"github.com/moguchev/telegram-bot/pkg/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := postgres.NewConnection(ctx, os.Getenv("DB_DSN"))
	if err != nil {
		logger.FatalKV("failed connect to database", "error", err)
	}

	bot, err := bot.New(
		os.Getenv("TELEGRAM_API_TOKEN"),
		bot.WithDebug(),
	)
	if err != nil {
		logger.FatalKV("failed create bot API", "error", err)
	}

	bot.Run()
}
