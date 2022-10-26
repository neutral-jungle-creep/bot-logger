package telegram

import (
	"bot_logger/internal/service"
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
	"bot_logger/internal/storage/pgSQL"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handler) AddUser(user *tgbotapi.User) error {
	conn, err := pgSQL.NewPgConnect()
	if err != nil {
		logrus.Infof("Ошибка подключения к базе данных: %s", err.Error())
		return err
	}
	defer conn.Close(context.Background())

	stor := storage.NewPgUserStorage(conn)
	serv := service.NewUserService(stor)
	userDto := dto.NewUserDto(user.ID, user.UserName, true)
	result := serv.AddUser(userDto)
	return result
}

func (h *Handler) EditUser(user *tgbotapi.User) error {
	conn, err := pgSQL.NewPgConnect()
	if err != nil {
		logrus.Infof("Ошибка подключения к базе данных: %s", err.Error())
		return err
	}
	defer conn.Close(context.Background())

	stor := storage.NewPgUserStorage(conn)
	serv := service.NewUserService(stor)
	userDto := dto.NewUserDto(user.ID, user.UserName, false)
	result := serv.EditUser(userDto)
	return result
}
