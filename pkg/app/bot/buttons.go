package bot

import "strconv"

type buttonText string

const (
	ButtonSettingsNotifications             = buttonText("🔔 Настройки уведомлений")
	ButtonSettingsNotificationsReviewsOn    = buttonText("🔔 Вкл. уведомления об отзывах")
	ButtonSettingsNotificationsReviewsOff   = buttonText("🔕 Выкл. уведомления об отзывах")
	ButtonSettingsNotificationsQuestionsOn  = buttonText("🔔 Вкл. уведомления о вопросах")
	ButtonSettingsNotificationsQuestionsOff = buttonText("🔕 Выкл. уведомления о вопросах")
)

var (
	CallbackText = map[callbackType]buttonText{
		SettingsNotifications: ButtonSettingsNotifications,
	}
)

func (ct callbackType) Data() string {
	return strconv.Itoa(int(ct))
}

func (ct callbackType) Text(args ...interface{}) string {
	switch ct {
	case SettingsNotificationsReviewsSwitch:
		if len(args) != 1 {
			return string(ButtonSettingsNotificationsReviewsOn)
		}
		st, ok := args[0].(bool)
		if !ok {
			return string(ButtonSettingsNotificationsReviewsOn)
		}
		if st {
			return string(ButtonSettingsNotificationsReviewsOff)
		}
		return string(ButtonSettingsNotificationsReviewsOn)
	case SettingsNotificationsQuestionsSwitch:
		if len(args) != 1 {
			return string(ButtonSettingsNotificationsQuestionsOn)
		}
		st, ok := args[0].(bool)
		if !ok {
			return string(ButtonSettingsNotificationsQuestionsOn)
		}
		if st {
			return string(ButtonSettingsNotificationsQuestionsOff)
		}
		return string(ButtonSettingsNotificationsQuestionsOn)
	default:
		return string(CallbackText[ct])
	}
}
