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
	userService, err := userComposite()
	if err != nil {
		return err
	}

	userDto := dto.NewUserDto(user.ID, user.UserName, true)
	result := userService.AddUser(userDto)
	return result
}

func (h *Handler) EditUser(user *tgbotapi.User) error {
	userService, err := userComposite()
	if err != nil {
		return err
	}

	userDto := dto.NewUserDto(user.ID, user.UserName, false)
	result := userService.EditUser(userDto)
	return result
}

func userComposite() (*service.UserService, error) {
	conn, err := pgSQL.NewPgConnect()
	if err != nil {
		logrus.Infof("Ошибка подключения к базе данных: %s", err.Error())
		return nil, err
	}

	stor := storage.NewPgUserStorage(conn)
	serv := service.NewUserService(stor)
	return &serv, nil
}
