package user

import (
	"errors"
	"strings"
	"unicode/utf8"
)

const (
	NAME_MAX_LENGTH_RUNE = 17
)

type Name struct {
	value string
}

func NewName(v string) (*Name, error) {
	if utf8.RuneCountInString(v) > NAME_MAX_LENGTH_RUNE {
		return nil, errors.New("名前文字列長が不正です。")
	}
	if strings.TrimSpace(v) == "" {
		return nil, errors.New("名前文字列が空白です。")
	}
	if strings.Contains(v, "  ") {
		return nil, errors.New("名前文字列に空白が連続して含まれています。")
	}
	if strings.HasPrefix(v, " ") {
		return nil, errors.New("名前文字列の先頭に空白が含まれています。")
	}
	if strings.HasSuffix(v, " ") {
		return nil, errors.New("名前文字列の末尾に空白が含まれています。")
	}

	return &Name{value: v}, nil
}

func (n *Name) NameValue() string {
	return n.value
}
