package telegram

import (
	"bot_logger/internal/service"
	"bot_logger/internal/service/dto"
	"bot_logger/internal/storage"
	"bot_logger/internal/storage/pgSQL"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func (h *Handler) AddMessage(message *tgbotapi.Message) error {
	messageService := messageComposite()
	messageDto := dto.NewMessageDto(message.MessageID, message.From.ID, parseTimeStamp(message.Date),
		message.Text, false)
	result := messageService.AddMessage(messageDto)
	return result
}

func (h *Handler) EditMessage(message *tgbotapi.Message) error {
	messageService := messageComposite()
	messageDto := dto.NewMessageDto(message.MessageID, message.From.ID, parseTimeStamp(message.Date),
		message.Text, true)
	result := messageService.EditMessage(messageDto)
	return result
}

func messageComposite() *service.MessageService {
	conn, err := pgSQL.NewPgConnect()
	if err != nil {
		logrus.Fatalf("Ошибка подключения к базе данных: %s", err.Error())
	}

	stor := storage.NewPgMessageStorage(conn)
	serv := service.NewMessageService(stor)
	return &serv
}

func parseTimeStamp(timeStamp int) string {
	tm, err := strconv.ParseInt(strconv.Itoa(timeStamp), 10, 64)
	if err != nil {
		return strconv.FormatInt(tm, 10)
	}

	ut := time.Unix(tm, 0)
	timeForStruct := ut.Format("2006-01-02T15:04:05")

	return timeForStruct
}
