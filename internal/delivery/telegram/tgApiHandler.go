package telegram

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/usecase"
	"bot_logger/pkg/exceptions"
	"bot_logger/pkg/logs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"time"
)

func Run(bot *tgbotapi.BotAPI, config *configs.Configuration) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// проверка на существование не записанных в базу данных логов
	unwrittenUpdate, err := exceptions.ReadUnwrittenUpdate(&config.UnwrittenDataFile)
	if err != nil {
		log.Println(logs.UnwrittenUpdateNil)
	} else {
		handleUpdate(bot, unwrittenUpdate, config)
		log.Println(logs.UnwrittenWasWrite)
		os.Remove("../unwritten_data.json")
	}

	// запуск запроса на поиск обновлений
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if checkChatAccess(update, config.AccessChatID) {
			handleUpdate(bot, update, config)
		} else {
			msg := exceptions.NewBotMessageForChat(bot, update.Message.Chat.ID, logs.NoAccess)
			msg.SendMessageToChat()
		}
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *configs.Configuration) {
	typeOfUpdate := defineUpdateType(&update)
	writeResult := typeOfUpdate.UpdateWrite(config)
	if writeResult != nil {
		except := exceptions.NewUpdateException(bot, &update, config)
		except.Run()
	}
}

func checkChatAccess(update tgbotapi.Update, chatID int64) bool {
	if update.FromChat().ID == chatID {
		return true
	} else {
		return false
	}
}

func defineUpdateType(update *tgbotapi.Update) UpdateWriter {
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
		Message: domain.NewMessage(args.id, args.date, false, args.text, args.messageSender),
	}
}

func NewEditMessage(m *tgbotapi.Message) *usecase.EditMessage {
	args := newArgsForMessage(m)

	return &usecase.EditMessage{
		Message: domain.NewMessage(args.id, args.date, true, args.text, args.messageSender),
	}
}

type UpdateWriter interface {
	UpdateWrite(config *configs.Configuration) error
}

type argsForMessage struct {
	id            string
	date          string
	v4Data        string
	text          string
	messageSender *domain.User
}

func newArgsForMessage(m *tgbotapi.Message) *argsForMessage {
	return &argsForMessage{
		id:            strconv.Itoa(m.MessageID),
		date:          parseTimeStamp(m.Date),
		text:          m.Text,
		messageSender: domain.NewUser(m.From.UserName, strconv.FormatInt(m.From.ID, 10), true),
	}
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
