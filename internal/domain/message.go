package domain

type Message struct {
	Id       int64
	SenderId int64
	Date     string
	Text     string
}

func NewMessage(id int64, senderId int64, date string, text string) *Message {
	return &Message{
		Id:       id,
		SenderId: senderId,
		Date:     date,
		Text:     text,
	}
}
