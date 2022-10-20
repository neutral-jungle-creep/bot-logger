package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/storage/pgSQL"
	"context"
)

type AddMessage struct {
	Message *domain.Message
}

type EditMessage struct {
	Message *domain.Message
}

func (a AddMessage) UpdateHandle(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	result := pgSQL.AddToDBMessage(a.Message, conn, config)
	conn.Close(context.Background())

	return result
}

func (e EditMessage) UpdateHandle(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	result := pgSQL.EditInDBMessage(e.Message, conn)
	conn.Close(context.Background())

	return result
}
