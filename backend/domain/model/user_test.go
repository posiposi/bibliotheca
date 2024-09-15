package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testUser *User
var userId UserId
var name Name
var email Email
var password Password

func TestMain(m *testing.M) {
	testUser = NewUser("user_1", "test_user", "test@email.com", "pass")
	userId = testUser.Id()
	name = testUser.Name()
	email = testUser.Email()
	password = testUser.Password()
	m.Run()
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		test_name string
		id        string
		name      string
		email     string
		password  string
		User      *User
	}{
		{"正常系", "user_1", "test_user", "test@email.com", "pass", testUser},
	}

	for _, test := range tests {
		t.Run(test.test_name, func(t *testing.T) {
			user := NewUser(test.id, test.name, test.email, test.password)
			assert.Equal(t, user, testUser, "NewUser is equal!")
		})
	}
}

func TestId(t *testing.T) {
	assert.IsType(t, userId, UserId{}, "Id is UserId type!")
	assert.Equal(t, userId.Value(), "user_1", "Id is equal!")
}

func TestName(t *testing.T) {
	assert.IsType(t, name, Name{}, "Name is Name type!")
	assert.Equal(t, name.Value(), "test_user", "Name is equal!")
}

func TestEmail(t *testing.T) {
	assert.IsType(t, email, Email{}, "Email is Email type!")
	assert.Equal(t, email.Value(), "test@email.com", "Email is equal!")
}

func TestPassword(t *testing.T) {
	assert.IsType(t, password, Password{}, "Password is Password type!")
	assert.Equal(t, password.Value(), "pass", "Password is equal!")
}
