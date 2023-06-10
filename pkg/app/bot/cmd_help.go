package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	helpCMD = `
	<b>Бот-помошник для селлеров Wildberries</b>
	💬 <b>Поддержка:</b> @LeoLeGrand, пишите если у вас возникли какие то проблемы.
	⌚ <b>Онлайн:</b> с 10:00 - 18:00 по мск.
	❗Ничего не покупаю и не беру на реализацию, рекламы в боте нет.`
)

func (b *bot) HelpCmd(upd tgbotapi.Update) {
	b.sendMessage(upd.Message.Chat.ID, helpCMD, true)
}
