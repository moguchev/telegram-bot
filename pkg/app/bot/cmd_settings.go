package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

func (b *bot) SettingsCmd(upd tgbotapi.Update) {

	ch, ok := b.chats.GetChat(ChatID(upd.Message.Chat.ID))
	if !ok {
		return
	}

	settings := ch.GetSecctings()

	message := "*Текущие настройки*:\n1. *Уведомления*"

	if settings.CommentPushesOn {
		message += "\nНовые комментрии: включены ✅"
	} else {
		message += "\nНовые комментрии: выключены ❌"
	}
	if settings.QuestionPushesOn {
		message += "\nНовые вопросы: включены ✅"
	} else {
		message += "\nНовые вопросы: выключены ❌"
	}

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
	reply.ParseMode = tgbotapi.ModeMarkdown

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(string(ButtonSettings), string(ButtonSettings)),
		),
	)
	reply.ReplyMarkup = keyboard

	if err := b.apiRequest(reply); err != nil {
		logger.Error("failed to send start message", zap.Error(err))
	}
}
