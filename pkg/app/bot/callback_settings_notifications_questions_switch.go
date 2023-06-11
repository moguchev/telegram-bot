package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *bot) SettingsNotificationsQuestionsSwitchCallback(upd tgbotapi.Update) {
	chatID := upd.CallbackQuery.Message.Chat.ID

	ch, ok := b.chats.GetChat(ChatID(chatID))
	if !ok {
		return
	}

	ch.SetQuestionsNotifications(!ch.GetSettings().QuestionsNotificationsOn)

	_ = b.apiRequest(tgbotapi.NewEditMessageReplyMarkup(
		chatID,
		upd.CallbackQuery.Message.MessageID,
		buildSettingsKeyboardMarkup(ch.GetSettings()),
	))
}
