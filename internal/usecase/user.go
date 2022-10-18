package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
)

func RunUser(user *domain.User, config configs.Configuration) {

}

func NewUser(name string, id string, active bool) *domain.User {
	return &domain.User{
		Username: name,
		UserId:   id,
		IsActive: active,
	}
}
