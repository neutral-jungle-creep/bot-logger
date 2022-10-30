package dto

type UserDto struct {
	Id       int64
	Username string
}

func NewUserDto(id int64, name string) *UserDto {
	return &UserDto{
		Id:       id,
		Username: name,
	}
}
