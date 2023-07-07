package bot

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type callbackType int

const (
	SettingsNotifications callbackType = iota + 1
	SettingsNotificationsReviewsSwitch
	SettingsNotificationsQuestionsSwitch
)

type callbackFn func(upd tgbotapi.Update)

func (b *bot) initCallbacks() {
	b.callbacks = map[callbackType]callbackFn{
		SettingsNotifications:                b.SettingsNotificationsCallback,
		SettingsNotificationsReviewsSwitch:   b.SettingsNotificationsReviewsSwitchCallback,
		SettingsNotificationsQuestionsSwitch: b.SettingsNotificationsQuestionsSwitchCallback,
	}
}

func (b *bot) HandleCallback(upd tgbotapi.Update) {
	data := upd.CallbackData()

	i, err := strconv.Atoi(data)
	if err != nil {
		return
	}

	cbFn, ok := b.callbacks[callbackType(i)]
	if !ok {
		return
	}

	cbFn(upd)
}
