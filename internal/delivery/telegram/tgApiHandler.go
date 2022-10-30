package telegram

import (
	"bot_logger/internal/storage/pgSQL"
	"bot_logger/pkg/exceptions"
	"bot_logger/pkg/logs"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Run(bot *tgbotapi.BotAPI) {

	// проверка на существование не записанных в базу данных обновлений
	unwrittenUpdate, err := exceptions.ReadUnwrittenUpdate(viper.GetString("unwrittenDataFile"))
	if err != nil {
		logrus.Println("Не записанных в базу данных сообщений не сущетсвует")
	} else {
		readNewUpdate(bot, unwrittenUpdate)
		logrus.Println("Обновления из файла были добавлены в базу данных")
		os.Remove(viper.GetString("unwrittenDataFile"))
	}

	// запуск запроса на поиск обновлений
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if checkChatAccess(update.FromChat().ID) {
			readNewUpdate(bot, &update)
		} else {
			msg := exceptions.NewBotMessageForChat(bot, update.Message.Chat.ID, logs.NoAccess)
			msg.SendMessageToChat()
		}
	}
}

func readNewUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	conn, err := pgSQL.NewPgConnect()
	if err != nil {
		logrus.Panicf("Ошибка подключения к базе данных: %s", err.Error())
	}
	defer conn.Close(context.Background())

	handler := HandlerComposite(conn)

	if err := updateHandle(update, handler); err != nil {
		exceptions.Run(bot, update)
	}
}

func updateHandle(update *tgbotapi.Update, h *Handler) error {
	if update.Message != nil {
		if update.Message.NewChatMembers != nil {
			return h.AddUser(&update.Message.NewChatMembers[0])
		}
		if update.Message.LeftChatMember != nil {
			return h.EditUser(update.Message.LeftChatMember)
		}
		return h.AddMessage(update.Message)
	} else if update.EditedMessage != nil {
		return h.EditMessage(update.EditedMessage)
	}
	return nil
}

func checkChatAccess(chatID int64) bool {
	if chatID == viper.GetInt64("accessChatID") {
		return true
	} else {
		return false
	}
}
