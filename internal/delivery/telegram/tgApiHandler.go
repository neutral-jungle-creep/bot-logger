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
			handleResult := typeOfUpdate.UpdateHandle(config)
			if handleResult != nil {
				msg := newBotMessageForChat(bot, update.Message.Chat.ID, logs.ErrWriteDB)
				msg.sendMessageToChat()
			}
		} else {
			log.Println("error code 403: no access")
			msg := newBotMessageForChat(bot, update.Message.Chat.ID, logs.NoAccess)
			msg.sendMessageToChat()
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
	UpdateHandle(config *configs.Configuration) error
}

type UpdateUser struct {
	User     *tgbotapi.User
	IsActive bool
}

type UpdateMessage struct {
	Message *tgbotapi.Message
	IsEdit  bool
}

func (u UpdateUser) UpdateHandle(config *configs.Configuration) error {
	user := usecase.NewUser(u.User.UserName, strconv.FormatInt(u.User.ID, 10), u.IsActive)
	usecaseResult := usecase.RunUser(user, config)

	return usecaseResult
}

func (u UpdateMessage) UpdateHandle(config *configs.Configuration) error {
	date := parseTimeStamp(u.Message.Date)
	messageSender := usecase.NewUser(u.Message.From.UserName, strconv.FormatInt(u.Message.From.ID, 10), true)
	messageText := newMessageText(u.Message.Text)

	message := usecase.NewMessage(strconv.Itoa(u.Message.MessageID), date, u.IsEdit, *messageSender, messageText)
	usecaseResult := usecase.RunMessage(message, config)

	return usecaseResult
}

func parseTimeStamp(timeStamp int) string {
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

type botMessageForChat struct {
	bot     *tgbotapi.BotAPI
	chatID  int64
	message string
}

func newBotMessageForChat(bot *tgbotapi.BotAPI, chatId int64, text string) *botMessageForChat {
	return &botMessageForChat{
		bot:     bot,
		chatID:  chatId,
		message: text,
	}
}

func (b botMessageForChat) sendMessageToChat() {
	msg := tgbotapi.NewMessage(b.chatID, b.message)
	b.bot.Send(msg)
}
