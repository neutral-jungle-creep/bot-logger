package storage

import (
	"bot_logger/internal/domain"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	addMessage = `INSERT INTO public.messages (message_id, date, text, is_edit, user_id) 
VALUES ($1, $2, $3, $4, (SELECT id FROM public.users WHERE tg_user_id = $5));`

	editMessage = `UPDATE public.messages SET text=$1, is_edit=$2 WHERE message_id=$3;`
)

type PgMessageStorage struct {
	conn *pgx.Conn
}

func NewPgMessageStorage(conn *pgx.Conn) *PgMessageStorage {
	return &PgMessageStorage{
		conn: conn,
	}
}

func (s *PgMessageStorage) AddMessage(message *domain.Message) error {
	_, err := s.conn.Exec(context.Background(), addMessage,
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

func (s *PgMessageStorage) EditMessage(message *domain.Message) error {
	_, err := s.conn.Exec(context.Background(), editMessage,
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
