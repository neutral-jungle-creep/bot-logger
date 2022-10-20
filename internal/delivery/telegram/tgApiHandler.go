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
				msg := newBotMessageForChat(bot, config.AdminTgChatID, logs.ErrWriteDB)
				msg.sendMessageToChat()
			}
		} else {
			log.Println("error code 403: no access")
			msg := newBotMessageForChat(bot, update.Message.Chat.ID, logs.NoAccess)
			msg.sendMessageToChat()
		}
	}
}

func checkChatAccess(update tgbotapi.Update, chatID int64) bool {
	if update.FromChat().ID == chatID {
		return true
	} else {
		return false
	}
}

func defineUpdateType(update *tgbotapi.Update) UpdateHandler {
	if update.Message != nil {
		if update.Message.NewChatMembers != nil {
			return NewAddUser(&update.Message.NewChatMembers[0])
		}
		if update.Message.LeftChatMember != nil {
			return NewEditUser(update.Message.LeftChatMember)
		}
		// обработка нового сообщения
		return NewAddMessage(update.Message)
	} else if update.EditedMessage != nil {
		// обработка отредактированного сообщения
		return NewEditMessage(update.EditedMessage)
	}
	return nil
}

type UpdateHandler interface {
	UpdateHandle(config *configs.Configuration) error
}

func NewAddUser(u *tgbotapi.User) *usecase.AddUser {
	return &usecase.AddUser{
		User: domain.NewUser(u.UserName, strconv.FormatInt(u.ID, 10), true),
	}
}

func NewEditUser(u *tgbotapi.User) *usecase.EditUser {
	return &usecase.EditUser{
		User: domain.NewUser(u.UserName, strconv.FormatInt(u.ID, 10), false),
	}
}

func NewAddMessage(m *tgbotapi.Message) *usecase.AddMessage {
	args := newArgsForMessage(m)

	return &usecase.AddMessage{
		Message: domain.NewMessage(args.id, args.date, false, args.messageSender, args.messageText),
	}
}

func NewEditMessage(m *tgbotapi.Message) *usecase.EditMessage {
	args := newArgsForMessage(m)

	return &usecase.EditMessage{
		Message: domain.NewMessage(args.id, args.date, true, args.messageSender, args.messageText),
	}
}

type argsForMessage struct {
	id            string
	date          string
	v4Data        string
	messageSender *domain.User
	messageText   *domain.MessageText
}

func newArgsForMessage(m *tgbotapi.Message) *argsForMessage {
	return &argsForMessage{
		id:            strconv.Itoa(m.MessageID),
		date:          parseTimeStamp(m.Date),
		messageSender: domain.NewUser(m.From.UserName, strconv.FormatInt(m.From.ID, 10), true),
		messageText:   newMessageText(m.Text),
	}
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

func newMessageText(messageText string) *domain.MessageText {
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

	return &text
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
