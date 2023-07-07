package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type commandEntity struct {
	key    commandKey
	desc   string
	action func(upd tgbotapi.Update)
}

type commandKey string

const (
	StartCmdKey    = commandKey("start")
	HelpCmdKey     = commandKey("help")
	SettingsCmdKey = commandKey("settings")
)

func (b *bot) initCommands() error {
	commands := []commandEntity{
		{
			key:    StartCmdKey,
			desc:   "▶️ Запустить бота",
			action: b.StartCmd,
		},
		{
			key:    HelpCmdKey,
			desc:   "🆘 Поддержка",
			action: b.HelpCmd,
		},
		{
			key:    SettingsCmdKey,
			desc:   "⚙️ Настройки",
			action: b.SettingsCmd,
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

func (b *bot) HandleCommand(upd tgbotapi.Update) {
	key := upd.Message.Command()

	if cmd, ok := b.commands[commandKey(key)]; ok {
		cmd.action(upd)
	} else {
		logger.Error("command handler not found", zap.String("cmd", key))
	}
}
