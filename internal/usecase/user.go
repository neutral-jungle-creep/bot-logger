package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/storage/pgSQL"
	"context"
)

type AddUser struct {
	User *domain.User
}

type EditUser struct {
	User *domain.User
}

func (u *AddUser) UpdateHandle(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	result := pgSQL.AddToDBUser(u.User, conn)
	conn.Close(context.Background())

	return result
}

func (u *EditUser) UpdateHandle(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	result := pgSQL.EditInDBUser(u.User, conn)
	conn.Close(context.Background())

	return result
}
