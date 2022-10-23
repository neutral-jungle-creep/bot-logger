package telegram

import (
	"bot_logger/internal/service"
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
	"bot_logger/internal/storage/pgSQL"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handler) AddUser(user *tgbotapi.User) error {
	userService := userComposite()
	userDto := dto.NewUserDto(user.ID, user.UserName, true)
	result := userService.AddUser(userDto)
	return result
}

func (h *Handler) EditUser(user *tgbotapi.User) error {
	userService := userComposite()
	userDto := dto.NewUserDto(user.ID, user.UserName, false)
	result := userService.EditUser(userDto)
	return result
}

func userComposite() *service.UserService {
	conn, err := pgSQL.NewPgConnect()
	if err != nil {
		logrus.Fatalf("Ошибка подключения к базе данных: %s", err.Error())
	}

	stor := storage.NewPgUserStorage(conn)
	serv := service.NewUserService(stor)
	return &serv
}
