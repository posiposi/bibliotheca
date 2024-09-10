package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	validEmail := strings.Repeat("a", 86) + "@example.co.jp"
	tests := []struct {
		test_name string
		value     string
		Email     *Email
		isErr     bool
	}{
		{"正常系", validEmail, &Email{value: validEmail}, false},
		{"正常系(最低限)", "a@b.co.jp", &Email{value: "a@b.co.jp"}, false},
		{"正常系(.com)", "a@b.com", &Email{value: "a@b.com"}, false},
		{"正常系(.ne.jp)", "a@b.ne.jp", &Email{value: "a@b.ne.jp"}, false},
		{"文字数オーバー", strings.Repeat("a", 87) + "@example.co.jp", nil, true},
		{"空白のみ", "  ", nil, true},
		{"空白が含まれる", "hoge @example.co.jp", nil, true},
		{"空文字", "", nil, true},
		{"メール形式不正", "hoge", nil, true},
		{"ドメイン不正", "hoge@hoge", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			email, err := NewEmail(tt.value)
			if (err != nil) != tt.isErr {
				t.Errorf("NewEmail() error = %v, isErr %v", err, tt.isErr)
				return
			}
			if err == nil {
				assert.Equal(t, email.value, tt.value, "NewEmail is equal!")
			}
		})
	}
}

func TestEmailValue(t *testing.T) {
	email, err := NewEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", email.EmailValue())
}
