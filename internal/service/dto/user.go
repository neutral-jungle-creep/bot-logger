package dto

type UserDto struct {
	Id       int64
	Username string
	IsActive bool
}

func NewUserDto(id int64, name string, isActive bool) *UserDto {
	return &UserDto{
		Id:       id,
		Username: name,
		IsActive: isActive,
	}
}
