package domain

type User struct {
	Username string
	UserId   string
	IsActive bool
}

func NewUser(name string, id string, active bool) *User {
	return &User{
		Username: name,
		UserId:   id,
		IsActive: active,
	}
}
