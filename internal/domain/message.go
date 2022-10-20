package domain

type Message struct {
	MessageId     string
	Date          string
	IsEdit        bool
	Text          string
	MessageSender User
}

func NewMessage(id string, date string, isEdit bool, text string, user *User) *Message {
	return &Message{
		MessageId:     id,
		Date:          date,
		IsEdit:        isEdit,
		Text:          text,
		MessageSender: *user,
	}
}
