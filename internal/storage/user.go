package storage

import (
	"bot_logger/internal/domain"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PgUserStorage struct {
	conn *pgx.Conn
}

func NewPgUserStorage(conn *pgx.Conn) *PgUserStorage {
	return &PgUserStorage{
		conn: conn,
	}
}

func (s *PgUserStorage) AddUserToDB(user *domain.User) error {
	_, err := s.conn.Exec(context.Background(), viper.GetString("queries.addUser"),
		user.Id,
		user.Username,
		user.IsActive,
	)

	if err != nil {
		logrus.Printf("Ошибка добавления пользователя в базу данных: %s", err.Error())
		return err
	}
	logrus.Printf("Добавление нового пользователя в базу данных: %v", *user)
	return nil
}

func (s *PgUserStorage) EditUserInDB(user *domain.User) error {
	_, err := s.conn.Exec(context.Background(), viper.GetString("queries.editUser"),
		user.IsActive,
		user.Id,
	)

	if err != nil {
		logrus.Printf("Ошибка редактирования пользователя в базе данных: %s", err.Error())
		return err
	}
	logrus.Printf("Редактирование пользователя в базе данных: %v", *user)
	return nil
}
