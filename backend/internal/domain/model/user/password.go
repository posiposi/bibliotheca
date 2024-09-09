package user

import (
	"errors"
	"strings"
)

const MIN_PASSWORD_LENGTH = 8
const MAX_PASSWORD_LENGTH = 64

type Password struct {
	value string
}

func NewPassword(v string) (*Password, error) {
	if len(v) < MIN_PASSWORD_LENGTH {
		return nil, errors.New("パスワード文字列長が不正です。")
	}
	if len(v) > MAX_PASSWORD_LENGTH {
		return nil, errors.New("パスワード文字列長が不正です。")
	}
	if strings.TrimSpace(v) == "" {
		return nil, errors.New("パスワード文字列が空白です。")
	}

	return &Password{value: v}, nil
}
