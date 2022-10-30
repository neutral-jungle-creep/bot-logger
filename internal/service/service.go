package service

import (
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
)

type User interface {
	AddUserToDB(u *dto.UserDto) error
	EditUserInDB(u *dto.UserDto) error
}

type Message interface {
	EditMessageInDB(m *dto.MessageDto) error
	AddMessageToDB(m *dto.MessageDto) error
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
