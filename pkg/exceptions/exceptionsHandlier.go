package exceptions

import (
	"bot_logger/configs"
	"bot_logger/pkg/logs"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	if err != nil {
		msg := NewBotMessageForChat(u.bot, u.config.AdminTgChatID, logs.ErrWriteUpdFile)
		msg.SendMessageToChat()
	} else {
		msg := NewBotMessageForChat(u.bot, u.config.AdminTgChatID, logs.ErrWriteDB)
		msg.SendMessageToChat()
	}
}

type botMessageForChat struct {
	bot     *tgbotapi.BotAPI
	chatID  int64
	message string
}

func NewBotMessageForChat(bot *tgbotapi.BotAPI, chatId int64, text string) *botMessageForChat {
	return &botMessageForChat{
		bot:     bot,
		chatID:  chatId,
		message: text,
	}
}

func (b botMessageForChat) SendMessageToChat() {
	msg := tgbotapi.NewMessage(b.chatID, b.message)
	b.bot.Send(msg)
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
