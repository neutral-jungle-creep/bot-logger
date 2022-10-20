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
	UnwrittenDataFile string `json:"unwrittenDataFile"`
	AccessChatID      int    `json:"accessChatID"`
	AdminTgChatID     int    `json:"adminTgChatID"`
	LinkToDB          string `json:"linkToDB"`
	Queries           struct {
		AddUser     string `json:"addUser"`
		EditUser    string `json:"editUser"`
		AddMessage  string `json:"addMessage"`
		EditMessage string `json:"editMessage"`
	} `json:"queries"`
}

func (c *Configuration) FillConfiguration(link string) {
	file, err := os.Open(link)
	if err != nil {
		log.Panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(data, &c); err != nil {
		log.Panic(err)
	}
}
