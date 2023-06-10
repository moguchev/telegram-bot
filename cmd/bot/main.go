package main

import (
	"os"

	"github.com/moguchev/telegram-bot/pkg/app/bot"
	"github.com/moguchev/telegram-bot/pkg/logger"
)

func main() {
	bot, err := bot.New(
		os.Getenv("TELEGRAM_API_TOKEN"),
		bot.WithDebug(),
	)
	if err != nil {
		logger.FatalKV("failed create bot API", "error", err)
	}

	bot.Run()
}
