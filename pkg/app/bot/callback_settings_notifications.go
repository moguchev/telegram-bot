package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *bot) SettingsNotificationsCallback(upd tgbotapi.Update) {
	chatID := upd.CallbackQuery.Message.Chat.ID

	ch, ok := b.chats.GetChat(ChatID(chatID))
	if !ok {
		return
	}

	reply := tgbotapi.NewMessage(chatID, "Настройка уведомлений:")
	reply.ParseMode = tgbotapi.ModeMarkdown
	reply.ReplyMarkup = buildSettingsNotificationsKeyboardMarkup(ch.GetSettings())

	_ = b.apiRequest(reply)
}

func buildSettingsNotificationsKeyboardMarkup(s Settings) tgbotapi.InlineKeyboardMarkup {
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
