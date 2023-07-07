package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/internal/app/bot/storage"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"gopkg.in/square/go-jose.v2/jwt"
)

const (
	invalidToken          = "Вы ввели неверный токен."
	tokenSuccesfullySaved = "Токен успешно сохранен."
	saveTokenFailed       = "Возникла ошибка при сохранении токена. Попробуйте еще раз."
)

func (b *bot) HandleText(upd tgbotapi.Update) {
	const api = "HandleText"
	var (
		chatID = upd.Message.Chat.ID
		id     = storage.ChatID(chatID)
	)

	ch, err := b.chats.GetChat(id)
	if err != nil {
		logger.ErrorKV(api,
			"action", "GetChat",
			"error", err,
		)
		return
	}

	if ch.GetToken() == "" {
		if _, err = jwt.ParseSigned(upd.Message.Text); err != nil {
			b.sendMessage(upd.Message.Chat.ID, invalidToken, false)
			return
		}

		if err = b.chats.SetToken(id, upd.Message.Text); err != nil {
			logger.ErrorKV(api,
				"action", "SetToken",
				"error", err,
			)
			b.sendMessage(upd.Message.Chat.ID, saveTokenFailed, false)
			return
		}
		b.sendMessage(upd.Message.Chat.ID, tokenSuccesfullySaved, false)
	}
}
