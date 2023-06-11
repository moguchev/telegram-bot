package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/app/bot/storage"
	"github.com/moguchev/telegram-bot/pkg/logger"
)

const (
	startMsg = `
<b>Что умеет бот?</b>
1. 🔔 Присылать уведомления о новых отзывах ⭐️ и вопросах ❓ в ваших карточках товаров.
`
	noAPIKeyMsg = `
<b>Отправьте боту x64 API ключ из вашего личного кабинета продавца.</b>

<b>Что такое API-Ключ и как его получить?</b>
API-ключ — секретный ключ, который Wildberries выдает своим поставщикам и используемый для получения данных с серверов Wildberries без доступа к вашему личному кабинету.
Найти его можно в разделе «<a href="https://seller.wildberries.ru/supplier-settings/access-to-api">Мой профиль – Доступ к API</a>». Это безопасно, все данные надежно зашифрованы и никому не доступны.

Доступ по API-ключу — это <b>самый безопасный способ получения данных</b>. Предоставляя нам Ваш API ключ мы получаем доступ только на получение тех данных, которые разрешает получить Wildberries (например: заказы, продажи, поступления, наличию на складах).`
)

func (b *bot) StartCmd(upd tgbotapi.Update) {
	const api = "StartCmd"

	b.sendMessage(upd.Message.Chat.ID, startMsg, true)

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
		b.sendMessage(upd.Message.Chat.ID, noAPIKeyMsg, true)
	}
}
