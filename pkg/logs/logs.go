package logs

import (
	"os"
)

const (
	NoAccess    = "Нет доступа"
	ConnDBError = "Ошибка подключения к базе данных"
	ConnCorrect = "Подключение к базе данных успешно"
)

// CreateOrOpenFileForLogs если файла не существует, создает его, если файл существует, открывает.
func CreateOrOpenFileForLogs(fileName *string) (*os.File, error) {

	file, err := os.OpenFile(*fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(*fileName)
		if err != nil {
			return file, err
		}
	}

	return file, nil
}
