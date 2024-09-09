package user

import (
	"errors"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

const (
	MIN_EMAIL_LENGTH = 3
	MAX_EMAIL_LENGTH = 100
)

func NewEmail(v string) (*Email, error) {
	if len(v) < MIN_EMAIL_LENGTH {
		return nil, errors.New("メールアドレスが短すぎます。")
	}
	if len(v) > MAX_EMAIL_LENGTH {
		return nil, errors.New("メールアドレスが長すぎます。")
	}
	if strings.ContainsAny(v, " ") {
		return nil, errors.New("メールアドレスに空白が含まれています。")
	}
	const emailRegex = `^[^\s@]+@[^\s@]+\.(co.jp|com|ne.jp)$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(v) {
		return nil, errors.New("メールアドレスの形式が不正です。")
	}

	return &Email{value: v}, nil
}

func (e *Email) EmailValue() string {
	return e.value
}
