package telegram

import (
	"bot_logger/internal/service/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

func (h *Handler) AddMessage(message *tgbotapi.Message) error {
	messageDto := dto.NewMessageDto(message.MessageID, message.From.ID, parseTimeStamp(message.Date),
		message.Text, false)
	result := h.service.AddChatMessage(messageDto)
	return result
}

func (h *Handler) EditMessage(message *tgbotapi.Message) error {
	messageDto := dto.NewMessageDto(message.MessageID, message.From.ID, parseTimeStamp(message.Date),
		message.Text, true)
	result := h.service.EditChatMessage(messageDto)
	return result
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
