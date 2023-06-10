package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/square/go-jose.v2/jwt"
)

const (
	invalidToken          = "Вы ввели неверный токен."
	tokenSuccesfullySaved = "Токен успешно сохранен."
)

func (b *bot) HandleTextCmd(upd tgbotapi.Update) {
	ch, ok := b.chats.GetChat(ChatID(upd.Message.Chat.ID))
	if !ok {
		return
	}

	if ch.GetToken() == "" {
		_, err := jwt.ParseSigned(upd.Message.Text)
		if err != nil {
			b.sendMessage(upd.Message.Chat.ID, invalidToken, false)
		} else {
			ch.SetToken(upd.Message.Text)

			reply := tgbotapi.NewMessage(upd.Message.Chat.ID, tokenSuccesfullySaved)

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
	}
}

type buttonValue string

const (
	ButtonSettings = buttonValue("🔔 Настройки уведомлений")
)
