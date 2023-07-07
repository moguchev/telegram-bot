package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/internal/app/bot/storage"
	"github.com/moguchev/telegram-bot/pkg/logger"
)

func (b *bot) SettingsNotificationsCallback(upd tgbotapi.Update) {
	const api = "SettingsNotificationsCallback"
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

	reply := tgbotapi.NewMessage(chatID, "Настройка уведомлений:")
	reply.ReplyMarkup = buildSettingsNotificationsKeyboardMarkup(ch.GetSettings())

	err = b.apiRequest(reply)
	if err != nil {
		logger.ErrorKV(api,
			"action", "apiRequest",
			"error", err,
		)
	}
}

func buildSettingsNotificationsKeyboardMarkup(s storage.ChatSettings) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				SettingsNotificationsReviewsSwitch.Text(s.ReviewsNotificationsOn),
				SettingsNotificationsReviewsSwitch.Data(),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				SettingsNotificationsQuestionsSwitch.Text(s.QuestionsNotificationsOn),
				SettingsNotificationsQuestionsSwitch.Data(),
			),
		),
	)
}
