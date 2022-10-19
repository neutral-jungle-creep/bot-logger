package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/storage/pgSQL"
	"context"
)

func RunUser(user *domain.User, config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	result := pgSQL.WriteUserToDB(user, conn)
	conn.Close(context.Background())

	return result
}

func NewUser(name string, id string, active bool) *domain.User {
	return &domain.User{
		Username: name,
		UserId:   id,
		IsActive: active,
	}
}
