package inmemmory

import (
	"sync"

	. "github.com/moguchev/telegram-bot/internal/app/bot/storage"
)

type chats struct {
	mx    sync.RWMutex
	chats map[ChatID]*chat
}

func NewChatsStorage() *chats {
	return &chats{
		chats: make(map[ChatID]*chat),
	}
}

func (chs *chats) AddChat(id ChatID, userInfo UserInfo) (Chat, error) {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	ch := NewChat(id, userInfo)
	chs.chats[id] = ch

	return ch, nil
}

func (chs *chats) DeleteChat(id ChatID) error {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	delete(chs.chats, id)
	return nil
}

func (chs *chats) GetChat(id ChatID) (Chat, error) {
	chs.mx.RLock()
	defer chs.mx.RUnlock()

	ch, ok := chs.chats[id]
	if !ok {
		return nil, ErrNotFound
	}

	return ch, nil
}

func (chs *chats) SetToken(id ChatID, token string) error {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	ch, ok := chs.chats[id]
	if !ok {
		return ErrNotFound
	}

	ch.token = token
	return nil
}

func (chs *chats) SetSettingsNotificationsQuestions(id ChatID, v bool) error {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	ch, ok := chs.chats[id]
	if !ok {
		return ErrNotFound
	}

	ch.ChatSettings.QuestionsNotificationsOn = v
	return nil
}

func (chs *chats) SetSettingsNotificationsReviews(id ChatID, v bool) error {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	ch, ok := chs.chats[id]
	if !ok {
		return ErrNotFound
	}

	ch.ChatSettings.ReviewsNotificationsOn = v
	return nil
}
