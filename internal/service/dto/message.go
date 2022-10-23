package dto

type MessageDto struct {
	Id       int
	SenderId int64
	Date     string
	Text     string
	IsEdit   bool
}

func NewMessageDto(id int, senderId int64, date string, text string, isEdit bool) *MessageDto {
	return &MessageDto{
		Id:       id,
		SenderId: senderId,
		Date:     date,
		Text:     text,
		IsEdit:   isEdit,
	}
}
