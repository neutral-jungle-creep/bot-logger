package domain

type Message struct {
	Id       int
	SenderId int64
	Date     string
	Text     string
	IsEdit   bool
}

func NewMessage(id int, senderId int64, date string, text string, isEdit bool) *Message {
	return &Message{
		Id:       id,
		SenderId: senderId,
		Date:     date,
		Text:     text,
		IsEdit:   isEdit,
	}
}
