package storage

import (
	"bot_logger/internal/domain"
	"github.com/jackc/pgx/v4"
)

type User interface {
	GetUser(user *domain.User) int
	AddUser(user *domain.User) error
	EditUser(user *domain.User) error
}

type Message interface {
	AddMessage(message *domain.Message) error
	EditMessage(message *domain.Message) error
}

type Storage struct {
	User
	Message
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		User:    NewPgUserStorage(conn),
		Message: NewPgMessageStorage(conn),
	}
}
