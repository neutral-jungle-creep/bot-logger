package logs

import (
	"os"
)

const (
	NoAccess    = "Нет доступа"
	ErrWriteDB  = "Произошла ошибка, сообщение не было записано в базу данных"
	ErrConnDB   = "Ошибка подключения к базе данных"
	ConnCorrect = "Подключение к базе данных успешно"
	ErrAddDBM   = "Ошибка добавления сообщения в базу данных: "
	ErrEditBDM  = "Ошибка редактирования сообщения в базе данных: "
	AddDBM      = "Добавление нового сообщения в базу данных: "
	EditDBM     = "Редактирование сообщения в базе данных: "
	ErrAddBDU   = "Ошибка добавления пользователя в базу данных: "
	ErrEditDBU  = "Ошибка редактирования пользователя в базе данных: "
	AddDBU      = "Добавление нового пользователя в базу данных: "
	EditBDU     = "Редактирование пользователя в базе данных: "
)

// CreateOrOpenFileForLogs если файла не существует, создает его, если файл существует, открывает.
func CreateOrOpenFileForLogs(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			return file, err
		}
	}

	return file, nil
}
