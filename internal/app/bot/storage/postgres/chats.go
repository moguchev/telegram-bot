package inmemmory

import (
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/moguchev/telegram-bot/internal/app/bot/storage"
)

type chats struct {
	pool *pgxpool.Pool
}

// NewChatsStorage returns chats storage
func NewChatsStorage(pool *pgxpool.Pool) *chats {
	return &chats{pool: pool}
}

/*

chats

id,deleted,created_at,updated_at


users

id, chat_id, first_name, last_name, created_at,updated_at

*/

func (chs *chats) AddChat(id ChatID, userInfo UserInfo) (Chat, error) {
	return nil, nil
}

func (chs *chats) DeleteChat(id ChatID) error {
	return nil
}

func (chs *chats) GetChat(id ChatID) (Chat, error) {
	return nil, nil
}

func (chs *chats) SetToken(id ChatID, token string) error {
	return nil
}

func (chs *chats) SetSettingsNotificationsQuestions(id ChatID, v bool) error {
	return nil
}

func (chs *chats) SetSettingsNotificationsReviews(id ChatID, v bool) error {
	return nil
}
