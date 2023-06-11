package bot

import "sync"

type ChatID int64

type Chats struct {
	mx    sync.RWMutex
	chats map[ChatID]Chat
}

type Settings struct {
	ReviewsNotificationsOn, QuestionsNotificationsOn bool
}

type Chat interface {
	GetToken() string
	SetToken(token string)
	GetSettings() Settings
	SetReviewsNotifications(v bool)
	SetQuestionsNotifications(v bool)
}

type chat struct {
	id    ChatID
	token string

	UserInfo
	Settings
}

type UserInfo struct {
	FromID    int64
	FirstName string
	LastName  string
}

func (c *chat) GetToken() string {
	return c.token
}

func (c *chat) SetToken(token string) {
	c.token = token
}

func (c *chat) GetSettings() Settings {
	return c.Settings
}

func (c *chat) SetReviewsNotifications(v bool) {
	c.Settings.ReviewsNotificationsOn = v
}

func (c *chat) SetQuestionsNotifications(v bool) {
	c.Settings.QuestionsNotificationsOn = v
}

func NewChats() *Chats {
	return &Chats{
		chats: make(map[ChatID]Chat),
	}
}

func (chs *Chats) AddChat(id ChatID, info UserInfo) Chat {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	ch := &chat{
		id:       id,
		UserInfo: info,
		Settings: Settings{},
	}
	chs.chats[id] = ch
	return ch
}

func (chs *Chats) DeleteChat(id ChatID) {
	chs.mx.Lock()
	defer chs.mx.Unlock()

	delete(chs.chats, id)
}

func (chs *Chats) GetChat(id ChatID) (Chat, bool) {
	chs.mx.RLock()
	defer chs.mx.RUnlock()

	ch, ok := chs.chats[id]
	return ch, ok
}
