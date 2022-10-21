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

func (u *AddUser) UpdateWrite(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	user := pgSQL.NewAddUser(u.User, conn, config)
	result := user.DBWrite()
	conn.Close(context.Background())

	return result
}

func (u *EditUser) UpdateWrite(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	user := pgSQL.NewEditUser(u.User, conn, config)
	result := user.DBWrite()
	conn.Close(context.Background())

	return result
}
