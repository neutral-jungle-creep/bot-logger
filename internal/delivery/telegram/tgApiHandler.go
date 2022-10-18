package telegram

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/usecase"
	"bot_logger/pkg/logs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

func Run(bot *tgbotapi.BotAPI, config *configs.Configuration) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// запуск запроса на поиск обновлений
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if checkChatAccess(update, config.AccessChatID) {
			typeOfUpdate := defineUpdateType(&update)
			typeOfUpdate.UpdateHandle(config)
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

func defineUpdateType(update *tgbotapi.Update) UpdateHandler {
	if update.Message != nil {
		if update.Message.NewChatMembers != nil {
			return UpdateUser{&update.Message.NewChatMembers[0], true}
		}
		if update.Message.LeftChatMember != nil {
			return UpdateUser{update.Message.LeftChatMember, false}
		}
		// обработка нового сообщения
		return UpdateMessage{update.Message, false}
	} else if update.EditedMessage != nil {
		// обработка отредактированного сообщения
		return UpdateMessage{update.EditedMessage, true}
	}
	return nil
}

type UpdateHandler interface {
	UpdateHandle(config *configs.Configuration)
}

type UpdateUser struct {
	User     *tgbotapi.User
	IsActive bool
}

type UpdateMessage struct {
	Message *tgbotapi.Message
	IsEdit  bool
}

func (u UpdateUser) UpdateHandle(config *configs.Configuration) {
	user := usecase.NewUser(u.User.UserName, strconv.FormatInt(u.User.ID, 10), u.IsActive)
	usecase.RunUser(user, config)
}

func (u UpdateMessage) UpdateHandle(config *configs.Configuration) {
	date := parseTime(u.Message.Date)
	messageSender := usecase.NewUser(u.Message.From.UserName, strconv.FormatInt(u.Message.From.ID, 10), true)
	messageText := newMessageText(u.Message.Text)

	message := usecase.NewMessage(strconv.Itoa(u.Message.MessageID), date, u.IsEdit, *messageSender, messageText)
	usecase.RunMessage(message, config)
}

func parseTime(timeStamp int) string {
	tm, err := strconv.ParseInt(strconv.Itoa(timeStamp), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	ut := time.Unix(tm, 0)
	timeForStruct := ut.Format("2006-01-02T15:04:05")

	return timeForStruct
}

func newMessageText(messageText string) domain.MessageText {
	var text = domain.MessageText{}

	for _, item := range strings.Split(messageText, "\n") {
		key, value := returnKeyAndValue(item, findSeparatorForString(&item))

		switch key {
		case "запрос":
			text.Query = value
		case "проблема":
			text.Problem = value
		case "причина":
			text.Cause = value
		case "решение":
			text.Solution = value
		case "источник":
			text.Source = value
		}
	}

	return text
}

// findSeparatorForString находит разделитель строки и возвращает его, возвращает 0, если в строке нет разделителя.
func findSeparatorForString(item *string) int {
	separator := strings.Index(*item, ":")
	if separator == -1 {
		separator = 0
	}

	return separator
}

// returnKeyAndValue принимает строку и индекс разделителя, делит строку по разделителю.
// Возвращает: key - часть строки до разделителя, value - часть строки после разделителя.
// Если строка является пустой, то возвращает две пустые строки, тем самым обрабатывает исключение:
// slice bounds out of range [1:0]
func returnKeyAndValue(text string, separator int) (string, string) {

	if text == "" {
		return "", ""
	}

	key := strings.ToLower(strings.TrimSpace(text[:separator]))
	value := strings.TrimSpace(text[separator+1:])

	return key, value
}

func sendMessageToChat(bot *tgbotapi.BotAPI, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, logs.NoAccess)
	bot.Send(msg)
}
