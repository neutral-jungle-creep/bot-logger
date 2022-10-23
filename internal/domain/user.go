package domain

type User struct {
	Id       int64
	Username string
	IsActive bool
}

func NewUser(id int64, name string, isActive bool) *User {
	return &User{
		Id:       id,
		Username: name,
		IsActive: isActive,
	}
}
