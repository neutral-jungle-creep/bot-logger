package service

import (
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
)

type User interface {
	AddChatMember(u *dto.UserDto) error
	LeaveChatMember(u *dto.UserDto) error
}

type Message interface {
	AddChatMessage(m *dto.MessageDto) error
	EditChatMessage(m *dto.MessageDto) error
}

type Service struct {
	User
	Message
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		User:    NewUserService(storage.User),
		Message: NewMessageService(storage.Message),
	}
}
