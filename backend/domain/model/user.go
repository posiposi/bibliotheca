package model

type User struct {
	id       UserId
	name     Name
	email    Email
	password Password
}

func NewUser(id, name, email, password string) *User {
	return &User{
		id:       UserId{value: id},
		name:     Name{value: name},
		email:    Email{value: email},
		password: Password{value: password},
	}
}

func (u *User) Id() UserId {
	return u.id
}

func (u *User) Name() Name {
	return u.name
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Password() Password {
	return u.password
}
