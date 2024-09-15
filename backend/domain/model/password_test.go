package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		test_name string
		value     string
		Password  *Password
		isErr     bool
	}{
		{"正常系(上限閾値)", strings.Repeat("a", 64), &Password{value: "password"}, false},
		{"正常系(下限閾値)", strings.Repeat("a", 8), &Password{value: "password"}, false},
		{"長過ぎる文字列", strings.Repeat("a", 65), nil, true},
		{"短すぎる文字列", strings.Repeat("a", 7), nil, true},
		{"文字列なし", "", nil, true},
		{"空文字", " ", nil, true},
	}

	for _, test := range tests {
		t.Run(test.test_name, func(t *testing.T) {
			password, err := NewPassword(test.value)
			if (err != nil) != test.isErr {
				t.Errorf("NewPassword() error = %v, isErr %v", err, test.isErr)
				return
			}
			if err == nil {
				assert.Equal(t, password.value, test.value, "NewPassword is equal!")
			}
		})
	}
}

func TestPasswordValue(t *testing.T) {
	password, err := NewPassword("testpassword")
	assert.NoError(t, err)
	assert.Equal(t, "testpassword", password.Value())
}
