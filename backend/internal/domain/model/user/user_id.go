package user

import (
	"errors"
	"strings"
)

const (
	USER_ID_MAX_LENGTH = 36
)

type UserId struct {
	value string
}

func NewUserId(id string) (*UserId, error) {
	if len(id) > USER_ID_MAX_LENGTH {
		return nil, errors.New("ユーザーID文字列長が不正です。")
	}
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("ユーザーID文字列が空白です。")
	}
	if strings.Contains(id, " ") {
		return nil, errors.New("ユーザーID文字列に空白が含まれています。")
	}

	userId := &UserId{
		value: id,
	}
	return userId, nil
}

func (u *UserId) UserIdValue() string {
	return u.value
}
