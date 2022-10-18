package telegram

import (
	"bot_logger/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func Run(bot *tgbotapi.BotAPI, config *configs.Configuration) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// запуск запроса на поиск обновлений
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if checkChatAccess(update, config.AccessChatID) {
			log.Println("ok")
		} else {
			log.Println("error code 403: no access")
			sendMessageToChat(bot, update.Message.Chat.ID)
		}
	}
}

func checkChatAccess(update tgbotapi.Update, chatID string) bool {
	if strconv.FormatInt(update.FromChat().ID, 10) == chatID {
		return true
	} else {
		return false
	}
}

func sendMessageToChat(bot *tgbotapi.BotAPI, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "Нет доступа")
	bot.Send(msg)
}
