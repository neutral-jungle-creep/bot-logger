package service

import (
	"bot_logger/internal/domain"
	"bot_logger/internal/service/dto"
)

type MessageStorage interface {
	AddMessageToDB(message *domain.Message) error
	EditMessageInDB(message *domain.Message) error
}

type MessageService struct {
	storage MessageStorage
}

func NewMessageService(storage MessageStorage) *MessageService {
	return &MessageService{
		storage: storage,
	}
}

func (s *MessageService) AddMessage(m *dto.MessageDto) error {
	message := domain.NewMessage(m.Id, m.SenderId, m.Date, m.Text, m.IsEdit)
	return s.storage.AddMessageToDB(message)
}

func (s *MessageService) EditMessage(m *dto.MessageDto) error {
	message := domain.NewMessage(m.Id, m.SenderId, m.Date, m.Text, m.IsEdit)
	return s.storage.EditMessageInDB(message)
}
