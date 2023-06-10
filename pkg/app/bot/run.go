package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"go.uber.org/zap"
)

// Run listens updates
func (b *bot) Run() {
	updatesCfg := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 10,
	}
	for upd := range b.GetUpdatesChan(updatesCfg) {
		go b.processUpdate(upd)
	}
}

func (b *bot) processUpdate(upd tgbotapi.Update) {
	if upd.MyChatMember != nil {
		// if user left or kicked bot
		switch upd.MyChatMember.NewChatMember.Status {
		case "left", "kicked":
			b.chats.DeleteChat(ChatID(upd.MyChatMember.Chat.ID))
		}
	}

	if upd.Message != nil {
		ch, exists := b.chats.GetChat(ChatID(upd.Message.Chat.ID))
		if !exists {
			ch = b.chats.AddChat(ChatID(upd.Message.Chat.ID), UserInfo{
				FromID:    upd.Message.From.ID,
				FirstName: upd.Message.From.FirstName,
				LastName:  upd.Message.From.LastName,
			})
		}
		_ = ch

		if upd.Message.IsCommand() {
			key := upd.Message.Command()

			if cmd, ok := b.commands[commandKey(key)]; ok {
				cmd.action(upd)
			} else {
				logger.Error("command handler not found", zap.String("cmd", key))
			}

			return
		} else {
			// TODO: another function
			b.HandleTextCmd(upd)
		}

		// if cmd, ok := b.replyToCommand(upd.Message.Text); ok {
		// 	cmd.action(upd)
		// 	return
		// }

		// if strings.HasPrefix(upd.Message.Text, onlinePrefix) {
		// 	b.OnlineCmd(upd)
		// } else if strings.HasPrefix(upd.Message.Text, promoPrefix) {
		// 	b.PromoCmd(upd)
		// } else if strings.HasPrefix(upd.Message.Text, userConfigPrefix) {
		// 	b.UpdateUserConfigCmd(upd)
		// } else {
		// 	b.SearchCmd(upd)
		// }
	}

	// if upd.CallbackQuery != nil {
	// 	data := upd.CallbackData()
	// 	entity := unmarshallCb(data)

	// 	if entity.cbType != Search && entity.parentType != Search {
	// 		b.clearSearchParams(upd.CallbackQuery.Message.Chat.ID)
	// 	}

	// 	callback := tgbotapi.NewCallback(upd.CallbackQuery.ID, "")
	// 	b.apiRequest(callback)

	// 	b.callbacks[entity.cbType](upd, entity)
	// }
}
