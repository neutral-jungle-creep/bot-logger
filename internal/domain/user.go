package domain

type User struct {
	Id       int64
	Username string
}

func NewUser(id int64, name string) *User {
	return &User{
		Id:       id,
		Username: name,
	}
}
