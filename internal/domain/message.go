package domain

type Message struct {
	Id       int
	SenderId int64
	Date     string
	Text     string
}

func NewMessage(id int, senderId int64, date string, text string) *Message {
	return &Message{
		Id:       id,
		SenderId: senderId,
		Date:     date,
		Text:     text,
	}
}
