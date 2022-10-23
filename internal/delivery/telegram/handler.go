package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TgUpdateHandler interface {
	AddUser(user *tgbotapi.User) error
	EditUser(user *tgbotapi.User) error
	AddMessage(message *tgbotapi.Message) error
	EditMessage(message *tgbotapi.Message) error
}

type Handler struct {
	handler TgUpdateHandler
}

func NewHandler(handler TgUpdateHandler) *Handler {
	return &Handler{
		handler: handler,
	}
}
