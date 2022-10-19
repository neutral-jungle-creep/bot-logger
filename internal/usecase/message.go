package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/storage/pgSQL"
	"context"
)

func RunMessage(message *domain.Message, config *configs.Configuration) error {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return err
	}

	result := pgSQL.WriteMessageToDB(message, conn)
	conn.Close(context.Background())

	return result
}

func NewMessage(id string, date string, isEdit bool, user domain.User, text domain.MessageText) *domain.Message {
	return &domain.Message{
		MessageId:     id,
		Date:          date,
		IsEdit:        isEdit,
		MessageSender: user,
		MessageText:   text,
	}
}
