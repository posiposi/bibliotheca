package model

type User struct {
	Id       UserId
	Name     Name
	Email    Email
	Password Password
}

func NewUser(id, name, email, password string) *User {
	return &User{
		Id:       UserId{value: id},
		Name:     Name{value: name},
		Email:    Email{value: email},
		Password: Password{value: password},
	}
}
