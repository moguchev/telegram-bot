package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *bot) SettingsNotificationsReviewsSwitchCallback(upd tgbotapi.Update) {
	chatID := upd.CallbackQuery.Message.Chat.ID

	ch, ok := b.chats.GetChat(ChatID(chatID))
	if !ok {
		return
	}

	ch.SetReviewsNotifications(!ch.GetSettings().ReviewsNotificationsOn)

	_ = b.apiRequest(tgbotapi.NewEditMessageReplyMarkup(
		chatID,
		upd.CallbackQuery.Message.MessageID,
		buildSettingsNotificationsKeyboardMarkup(ch.GetSettings()),
	))
}
