package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type commandEntity struct {
	key    commandKey
	desc   string
	action func(upd tgbotapi.Update)
}

type commandKey string

const (
	StartCmdKey = commandKey("start")
	HelpCmdKey  = commandKey("help")
)

func (b *bot) initCommands() error {
	commands := []commandEntity{
		{
			key:    StartCmdKey,
			desc:   "Запустить бота",
			action: b.StartCmd,
		},
		{
			key:    HelpCmdKey,
			desc:   "Поддержка",
			action: b.HelpCmd,
		},
	}

	tgCommands := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		b.commands[cmd.key] = cmd

		tgCommands = append(tgCommands, tgbotapi.BotCommand{
			Command:     "/" + string(cmd.key),
			Description: cmd.desc,
		})
	}

	config := tgbotapi.NewSetMyCommands(tgCommands...)
	return b.apiRequest(config)
}
