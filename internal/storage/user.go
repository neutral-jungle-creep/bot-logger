package storage

import (
	"bot_logger/internal/domain"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	addUser = `INSERT INTO public.users (tg_user_id, tg_user_name, active_user) VALUES ($1, $2, $3);`

	editUser = `UPDATE public.users SET active_user=$1 WHERE tg_user_id=$2;`

	getUser = `SELECT id FROM public.users WHERE tg_user_id=$1`
)

type PgUserStorage struct {
	conn *pgx.Conn
}

func NewPgUserStorage(conn *pgx.Conn) *PgUserStorage {
	return &PgUserStorage{
		conn: conn,
	}
}

func (s *PgUserStorage) GetUser(user *domain.User) int {
	var userId int
	result := s.conn.QueryRow(context.Background(), getUser, user.Id)
	if err := result.Scan(&userId); err != nil {
		return 0
	}
	return userId
}

func (s *PgUserStorage) AddUser(user *domain.User) error {
	_, err := s.conn.Exec(context.Background(), addUser,
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

func (s *PgUserStorage) EditUser(user *domain.User) error {
	_, err := s.conn.Exec(context.Background(), editUser,
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
