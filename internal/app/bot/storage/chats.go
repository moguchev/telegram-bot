package storage

import "errors"

type (
	ChatID int64

	ChatSettings struct {
		ReviewsNotificationsOn   bool
		QuestionsNotificationsOn bool
	}

	UserInfo struct {
		ID        int64
		FirstName string
		LastName  string
	}
)

type Chat interface {
	GetToken() string
	GetSettings() ChatSettings
	GetUserInfo() UserInfo
}

type ChatsStorage interface {
	GetChat(id ChatID) (Chat, error)
	AddChat(id ChatID, userInfo UserInfo) (Chat, error)
	DeleteChat(id ChatID) error

	SetToken(id ChatID, token string) error
	SetSettingsNotificationsQuestions(id ChatID, v bool) error
	SetSettingsNotificationsReviews(id ChatID, v bool) error
}

var (
	ErrNotFound = errors.New("not found")
)
