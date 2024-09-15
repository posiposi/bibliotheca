package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserId(t *testing.T) {
	tests := []struct {
		test_name string
		id        string
		UserId    *UserId
		isErr     bool
	}{
		{"正常系", "okuserlnbzuhefwjcwrszartdfyxnywyurlm", &UserId{value: "okuserlnbzuhefwjcwrszartdfyxnywyurlm"}, false},
		{"文字数オーバー", "nguserlnbzuhefwjcwrszartdfyxnywyurlmo", nil, true},
		{"空文字含む(前後)", " 12345 67890 ", nil, true},
		{"空文字含む(前)", " 1234567890", nil, true},
		{"空文字含む(後)", "1234567890 ", nil, true},
		{"文字列なし", "", nil, true},
		{"空文字", " ", nil, true},
	}

	for _, test := range tests {
		t.Run(test.test_name, func(t *testing.T) {
			id, err := NewUserId(test.id)
			if (err != nil) != test.isErr {
				t.Errorf("NewUserId() error = %v, isErr %v", err, test.isErr)
				return
			}
			if err == nil {
				assert.Equal(t, id.value, test.id, "NewUserId is equal!")
			}
		})
	}
}

func TestUserIdValue(t *testing.T) {
	id, err := NewUserId("testuser_id")
	assert.NoError(t, err)
	assert.Equal(t, "testuser_id", id.Value())
}
