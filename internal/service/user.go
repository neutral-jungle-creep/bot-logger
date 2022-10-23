package service

import (
	"bot_logger/internal/domain"
	"bot_logger/internal/service/dto"
)

type UserStorage interface {
	AddUser(user *domain.User) error
	EditUser(user *domain.User) error
}

type UserService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) UserService {
	return UserService{
		storage: storage,
	}
}

func (s *UserService) AddUser(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, u.IsActive)
	return s.storage.AddUser(user)
}

func (s *UserService) EditUser(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, u.IsActive)
	return s.storage.EditUser(user)
}
