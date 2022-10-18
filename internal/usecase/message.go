package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
)

func RunMessage(message *domain.Message, config configs.Configuration) {

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
