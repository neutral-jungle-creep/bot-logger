package configs

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Configuration struct {
	Token             string `json:"token"`
	LogFile           string `json:"logFile"`
	LinkToDB          string `json:"linkToDB"`
	UnwrittenDataFile string `json:"unwrittenDataFile"`
	AccessChatID      string `json:"accessChatID"`
}

func NewConfigFromFile(link string) *Configuration {
	var config Configuration

	file, err := os.Open(link)
	if err != nil {
		log.Panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Panic(err)
	}

	return &config
}
