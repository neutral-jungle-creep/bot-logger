package telegram

import (
	"bot_logger/internal/service/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) AddUser(user *tgbotapi.User) error {
	userDto := dto.NewUserDto(user.ID, user.UserName)
	result := h.service.AddChatMember(userDto)
	return result
}

func (h *Handler) EditUser(user *tgbotapi.User) error {
	userDto := dto.NewUserDto(user.ID, user.UserName)
	result := h.service.LeaveChatMember(userDto)
	return result
}
