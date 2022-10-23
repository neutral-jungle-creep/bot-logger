package storage

import (
	"bot_logger/internal/domain"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PgMessageStorage struct {
	db *pgx.Conn
}

func NewPgMessageStorage(db *pgx.Conn) *PgMessageStorage {
	return &PgMessageStorage{
		db: db,
	}
}

func (s *PgUserStorage) AddMessageToDB(message *domain.Message) error {
	_, err := s.conn.Exec(context.Background(), viper.GetString("queries.addMessage"),
		message.Id,
		message.Date,
		message.Text,
		false,
		message.SenderId,
	)

	if err != nil {
		logrus.Printf("Ошибка добавления сообщения в базу данных: %s", err.Error())
		return err
	}
	logrus.Printf("Добавление нового сообщения в базу данных: %v", *message)
	return nil
}

func (s *PgUserStorage) EditMessageInDB(message *domain.Message) error {
	_, err := s.conn.Exec(context.Background(), viper.GetString("queries.editMessage"),
		message.Text,
		true,
		message.Id,
	)

	if err != nil {
		logrus.Printf("Ошибка редактирования сообщения в базе данных: %s", err.Error())
		return err
	}
	logrus.Printf("Редактирование сообщения в базе данных: %v", *message)
	return nil
}
