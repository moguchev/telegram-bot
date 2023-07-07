package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/internal/app/bot/storage"
	"github.com/moguchev/telegram-bot/pkg/logger"
)

func (b *bot) SettingsNotificationsReviewsSwitchCallback(upd tgbotapi.Update) {
	const api = "SettingsNotificationsReviewsSwitchCallback"
	var (
		chatID = upd.CallbackQuery.Message.Chat.ID
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

	err = b.chats.SetSettingsNotificationsReviews(id,
		!ch.GetSettings().ReviewsNotificationsOn,
	)
	if err != nil {
		logger.ErrorKV(api,
			"action", "SetSettingsNotificationsReviews",
			"error", err,
		)
		return
	}

	err = b.apiRequest(tgbotapi.NewEditMessageReplyMarkup(
		chatID,
		upd.CallbackQuery.Message.MessageID,
		buildSettingsNotificationsKeyboardMarkup(ch.GetSettings()),
	))
	if err != nil {
		logger.ErrorKV(api,
			"action", "apiRequest",
			"error", err,
		)
	}
}
