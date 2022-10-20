package main

import (
	"bot_logger/configs"
	"bot_logger/internal/delivery/telegram"
	"bot_logger/pkg/logs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"os"
)

func main() {
	var config configs.Configuration
	config.FillConfiguration("../configs/config.json")

	// запись логов в файл
	file, err := logs.CreateOrOpenFileForLogs(&config.LogFile)
	if err != nil {
		log.Panic("Error in file for logs:", err)
	}
	wrt := io.MultiWriter(os.Stdout, file)
	log.SetOutput(wrt)

	defer file.Close()

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic("Telegram bot error:", err)
	}

	bot.Debug = false
	log.Println("Bot has been started!", "Bot name is:", bot.Self.UserName)

	telegram.Run(bot, &config)
}
