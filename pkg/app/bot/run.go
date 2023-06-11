package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		_, exists := b.chats.GetChat(ChatID(upd.Message.Chat.ID))
		if !exists {
			_ = b.chats.AddChat(ChatID(upd.Message.Chat.ID), UserInfo{
				FromID:    upd.Message.From.ID,
				FirstName: upd.Message.From.FirstName,
				LastName:  upd.Message.From.LastName,
			})
		}

		if upd.Message.IsCommand() {
			b.HandleCommand(upd)
		} else {
			b.HandleText(upd)
		}
	}

	if upd.CallbackQuery != nil {
		b.HandleCallback(upd)
	}
}
