package bot

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/app/bot/storage"
	"github.com/moguchev/telegram-bot/pkg/logger"
)

// Run listens updates
func (b *bot) Run() {
	updatesCfg := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 10,
	}
	for upd := range b.GetUpdatesChan(updatesCfg) {
		go b.processUpdate(upd)
	}
}

func (b *bot) processUpdate(upd tgbotapi.Update) {
	const api = "processUpdate"

	if upd.MyChatMember != nil {
		// if user left or kicked bot
		switch upd.MyChatMember.NewChatMember.Status {
		case "left", "kicked":
			_ = b.chats.DeleteChat(storage.ChatID(upd.MyChatMember.Chat.ID))
		}
	}

	if upd.Message != nil {
		id := storage.ChatID(upd.Message.Chat.ID)
		if _, err := b.chats.GetChat(id); errors.Is(err, storage.ErrNotFound) {
			if _, err = b.chats.AddChat(id, storage.UserInfo{
				ID:        upd.Message.From.ID,
				FirstName: upd.Message.From.FirstName,
				LastName:  upd.Message.From.LastName,
			}); err != nil {
				logger.ErrorKV(api,
					"action", "AddChat",
					"error", err,
				)
				return
			}
		}

		if upd.Message.IsCommand() {
			b.HandleCommand(upd)
		} else {
			b.HandleText(upd)
		}
	}

	if upd.CallbackQuery != nil {
		b.HandleCallback(upd)
	}
}
