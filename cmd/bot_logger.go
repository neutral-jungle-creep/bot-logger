package main

import (
	"bot_logger/configs"
	"bot_logger/internal/delivery/telegram"
	"bot_logger/pkg/exceptions"
	"bot_logger/pkg/logs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func main() {
	if err := configs.InitConfig("../configs"); err != nil {
		logrus.Fatalf("init config error: %s", err.Error())
	}

	// запись логов в файл
	file, err := logs.CreateOrOpenFileForLogs(viper.GetString("logFile"))
	if err != nil {
		logrus.Fatalf("Error in file for logs: %s", err.Error())
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, file))

	bot, err := tgbotapi.NewBotAPI(viper.GetString("token"))
	if err != nil {
		logrus.Fatalf("Telegram bot error: %s", err.Error())
	}

	bot.Debug = false
	logrus.Printf("Bot has been started! Bot name is: %s", bot.Self.UserName)

	defer file.Close()
	defer exceptions.NewBotMessageForChat(bot, viper.GetInt64("adminsTgChatID"), logs.FatalErrBot).SendMessageToChat()

	telegram.Run(bot)
}
