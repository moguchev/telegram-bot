package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

type bot struct {
	*tgbotapi.BotAPI

	commands map[commandKey]commandEntity
	chats    *Chats
}

// New creates bot instance
func New(token string, opts ...Option) (*bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	var o options
	for _, opt := range opts {
		opt(&o)
	}

	api.Debug = o.debug

	b := &bot{
		BotAPI:   api,
		commands: make(map[commandKey]commandEntity),
		chats:    NewChats(),
	}

	if err := b.initCommands(); err != nil {
		return nil, err
	}

	logger.Info("bot created", zap.String("username", api.Self.UserName))
	return b, nil
}

func (b *bot) apiRequest(c tgbotapi.Chattable) error {
	_, err := b.Request(c)
	return err
}

func (b *bot) sendMessage(chatID int64, message string, isHTML bool) {
	msg := tgbotapi.NewMessage(chatID, message)
	if isHTML {
		msg.ParseMode = tgbotapi.ModeHTML
	}

	if err := b.apiRequest(msg); err != nil {
		logger.Error("failed to send help message", zap.Error(err))
	}
}
