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

func (s *UserService) AddUserToDB(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, u.IsActive)
	return s.storage.AddUser(user)
}

func (s *UserService) EditUserInDB(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, u.IsActive)
	return s.storage.EditUser(user)
}
