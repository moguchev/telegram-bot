package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/square/go-jose.v2/jwt"
)

const (
	invalidToken          = "Вы ввели неверный токен."
	tokenSuccesfullySaved = "Токен успешно сохранен."
)

func (b *bot) HandleText(upd tgbotapi.Update) {
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

			b.sendMessage(upd.Message.Chat.ID, tokenSuccesfullySaved, false)
		}
	}
}
