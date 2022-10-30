package service

import (
	"bot_logger/internal/domain"
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
)

type MessageService struct {
	storage storage.Message
}

func NewMessageService(storage storage.Message) *MessageService {
	return &MessageService{
		storage: storage,
	}
}

func (s *MessageService) AddMessageToDB(m *dto.MessageDto) error {
	message := domain.NewMessage(m.Id, m.SenderId, m.Date, m.Text, m.IsEdit)
	return s.storage.AddMessage(message)
}

func (s *MessageService) EditMessageInDB(m *dto.MessageDto) error {
	message := domain.NewMessage(m.Id, m.SenderId, m.Date, m.Text, m.IsEdit)
	return s.storage.EditMessage(message)
}
