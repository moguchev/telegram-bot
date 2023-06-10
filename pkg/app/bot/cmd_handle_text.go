package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/moguchev/telegram-bot/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/square/go-jose.v2/jwt"
)

const (
	invalidToken          = "Вы ввели неверный токен."
	tokenSuccesfullySaved = "Токен успешно сохранен."
)

func (b *bot) HandleTextCmd(upd tgbotapi.Update) {
	ch, ok := b.chats.GetChat(ChatID(upd.Message.Chat.ID))
	if !ok {
		return
	}

	if ch.GetToken() == "" {
		_, err := jwt.ParseSigned(upd.Message.Text)
		if err != nil {
			b.sendMessage(upd.Message.Chat.ID, invalidToken, false)
		} else {
			ch.SetToken(upd.Message.Text)

			reply := tgbotapi.NewMessage(upd.Message.Chat.ID, tokenSuccesfullySaved)
			// reply.ParseMode = "html"

			var (
				commentsButton, questionButton replyKeyboardValue
				settings                       = ch.GetSecctings()
			)

			if settings.CommentPushesOn {
				commentsButton = ReplyCommentPushesOff
			} else {
				commentsButton = ReplyCommentPushesOn
			}

			if settings.QuestionPushesOn {
				questionButton = ReplyQuestionsPushesOff
			} else {
				questionButton = ReplyQuestionsPushesOn
			}

			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(string(commentsButton), string(commentsButton)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(string(questionButton), string(questionButton)),
				),
			)
			reply.ReplyMarkup = keyboard

			if err := b.apiRequest(reply); err != nil {
				logger.Error("failed to send start message", zap.Error(err))
			}
		}
	}
}

type replyKeyboardValue string

const (
	ReplyCommentPushesOn    = replyKeyboardValue("Уведомления: комментарии ✅")
	ReplyQuestionsPushesOn  = replyKeyboardValue("Уведомления: вопросы ✅")
	ReplyCommentPushesOff   = replyKeyboardValue("Уведомления: комментарии ❌")
	ReplyQuestionsPushesOff = replyKeyboardValue("Уведомления: вопросы ❌")
	// ✅ Включить оповещение о новых комментариях

)
