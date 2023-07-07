package inmemmory

import . "github.com/moguchev/telegram-bot/internal/app/bot/storage"

type chat struct {
	id    ChatID
	token string

	UserInfo
	ChatSettings
}

func NewChat(id ChatID, userInfo UserInfo) *chat {
	return &chat{
		id:       id,
		UserInfo: userInfo,
	}
}

func (c *chat) GetToken() string {
	return c.token
}

func (c *chat) GetUserInfo() UserInfo {
	return c.UserInfo
}

func (c *chat) GetSettings() ChatSettings {
	return c.ChatSettings
}
