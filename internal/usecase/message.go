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

func (a AddMessage) UpdateWrite(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	message := pgSQL.NewAddMessage(a.Message, conn, config)
	result := message.DBWrite()
	conn.Close(context.Background())

	return result
}

func (e EditMessage) UpdateWrite(config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	message := pgSQL.NewEditMessage(e.Message, conn, config)
	result := message.DBWrite()
	conn.Close(context.Background())

	return result
}
