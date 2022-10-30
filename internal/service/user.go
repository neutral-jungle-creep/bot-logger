package service

import (
	"bot_logger/internal/domain"
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
)

type UserService struct {
	storage storage.User
}

func NewUserService(storage storage.User) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) AddChatMember(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, true)
	if s.storage.GetUser(user) != 0 {
		return s.storage.EditUser(user)
	} else {
		return s.storage.AddUser(user)
	}
}

func (s *UserService) LeaveChatMember(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, false)
	return s.storage.EditUser(user)
}
