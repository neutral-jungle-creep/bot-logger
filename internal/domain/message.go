package domain

type Message struct {
	MessageId     string
	Date          string
	IsEdit        bool
	V4Data        string
	MessageSender *User
	MessageText   *MessageText
}

type MessageText struct {
	Query    string
	Problem  string
	Cause    string
	Solution string
	Source   string
}

func NewMessage(id string, date string, isEdit bool, v4Data string, user *User, text *MessageText) *Message {
	return &Message{
		MessageId:     id,
		Date:          date,
		IsEdit:        isEdit,
		V4Data:        v4Data,
		MessageSender: user,
		MessageText:   text,
	}
}
