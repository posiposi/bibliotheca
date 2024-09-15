package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		test_name string
		id        string
		name      string
		email     string
		password  string
		User      *User
	}{
		{"正常系", "user_1", "test_user", "test@email.com", "pass", &User{Id: UserId{value: "user_1"}, Name: Name{value: "test_user"}, Email: Email{value: "test@email.com"}, Password: Password{value: "pass"}}},
	}

	for _, test := range tests {
		t.Run(test.test_name, func(t *testing.T) {
			user := NewUser(test.id, test.name, test.email, test.password)
			assert.Equal(t, user, test.User, "NewUser is equal!")
		})
	}
}
