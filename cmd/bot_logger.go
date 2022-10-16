package main

import (
	"bot_logger/configs"
	"bot_logger/pkg/logs"
	"io"
	"log"
	"os"
)

func main() {
	var config = configs.NewConfigFromFile("../configs/config.json")

	// запись логов в файл
	file, err := logs.CreateOrOpenFileForLogs(&config.LogFile)
	if err != nil {
		log.Panic("Error in file for logs:", err)
	}
	wrt := io.MultiWriter(os.Stdout, file)
	log.SetOutput(wrt)

	defer file.Close()
}
