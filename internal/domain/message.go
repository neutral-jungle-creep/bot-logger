package domain

type Message struct {
	MessageId     string
	Date          string
	IsEdit        bool
	MessageSender User
	MessageText   MessageText
}

type MessageText struct {
	Query    string
	Problem  string
	Cause    string
	Solution string
	Source   string
}

func NewMessage(id string, date string, isEdit bool, user *User, text *MessageText) *Message {
	return &Message{
		MessageId:     id,
		Date:          date,
		IsEdit:        isEdit,
		MessageSender: *user,
		MessageText:   *text,
	}
}
