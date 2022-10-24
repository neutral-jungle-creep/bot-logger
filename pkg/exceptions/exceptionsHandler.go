package exceptions

import (
	"bot_logger/pkg/logs"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func ReadUnwrittenUpdate(fileLink string) (*tgbotapi.Update, error) {
	var update tgbotapi.Update

	file, err := os.Open(fileLink)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	textInFile, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(textInFile, &update); err != nil {
		return nil, err
	}

	return &update, nil
}

func Run(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	err := createFileWriteUpdate(update)
	if err == nil {
		msg := NewBotMessageForChat(bot, viper.GetInt64("adminsTgChatID"), logs.ErrBot)
		msg.SendMessageToChat()
	}
	logrus.Panic()
}

func createFileWriteUpdate(update *tgbotapi.Update) error {
	file, err := os.Create(viper.GetString("unwrittenDataFile"))
	if err != nil {
		return err
	}
	defer file.Close()

	jsonUpdate, err := json.Marshal(update)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonUpdate))
	if err != nil {
		return err
	}

	return nil
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
