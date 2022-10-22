package exceptions

import (
	"bot_logger/configs"
	"bot_logger/pkg/logs"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"os"
)

type updateException struct {
	bot    *tgbotapi.BotAPI
	update *tgbotapi.Update
	config *configs.Configuration
}

func NewUpdateException(bot *tgbotapi.BotAPI, update *tgbotapi.Update, config *configs.Configuration) *updateException {
	return &updateException{
		bot:    bot,
		update: update,
		config: config,
	}
}

func (u *updateException) Run() {
	err := u.createFileWriteUpdate()
	if err == nil {
		msg := NewBotMessageForChat(u.bot, u.config.AdminsTgChatID, logs.ErrWriteDB)
		msg.SendMessageToChat()
	}
	log.Panic()
}

type botMessageForChat struct {
	bot     *tgbotapi.BotAPI
	chatID  []int64
	message string
}

func NewBotMessageForChat(bot *tgbotapi.BotAPI, chatId []int64, text string) *botMessageForChat {
	return &botMessageForChat{
		bot:     bot,
		chatID:  chatId,
		message: text,
	}
}

func (b botMessageForChat) SendMessageToChat() {
	for _, id := range b.chatID {
		msg := tgbotapi.NewMessage(id, b.message)
		b.bot.Send(msg)
	}
}

func (u *updateException) createFileWriteUpdate() error {
	file, err := os.Create(u.config.UnwrittenDataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonUpdate, err := json.Marshal(u.update)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonUpdate))
	if err != nil {
		return err
	}

	return nil
}

func ReadUnwrittenUpdate(fileName *string) (tgbotapi.Update, error) {
	var update tgbotapi.Update

	file, err := os.Open(*fileName)
	if err != nil {
		return update, err
	}
	defer file.Close()

	textInFile, err := io.ReadAll(file)
	if err != nil {
		return update, err
	}

	if err := json.Unmarshal(textInFile, &update); err != nil {
		return update, err
	}

	return update, nil
}
