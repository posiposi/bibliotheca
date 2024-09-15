package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		test_name string
		value     string
		Name      *Name
		isErr     bool
	}{
		{"正常系", "正常系テストユーザー", &Name{value: "okuserlnbzuhefwjcwrszartdfyxnywyurlm"}, false},
		{"長過ぎる文字列", "nguserlnbzuhefwjcwrszartdfyxnywyurlmo", nil, true},
		{"先頭が空文字", " 1234567890", nil, true},
		{"末尾が空文字", "1234567890 ", nil, true},
		{"文字列なし", "", nil, true},
		{"空文字", " ", nil, true},
	}

	for _, test := range tests {
		t.Run(test.test_name, func(t *testing.T) {
			name, err := NewName(test.value)
			if (err != nil) != test.isErr {
				t.Errorf("NewName() error = %v, isErr %v", err, test.isErr)
				return
			}
			if err == nil {
				assert.Equal(t, name.value, test.value, "NewName is equal!")
			}
		})
	}
}

func TestNameValue(t *testing.T) {
	name, err := NewName("testuser")
	assert.NoError(t, err)
	assert.Equal(t, "testuser", name.Value())
}
