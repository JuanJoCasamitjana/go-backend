package model

import (
	"errors"
	"unicode/utf8"
)

const (
	MaxTitleLength       = 128
	MaxDescriptionLength = 1024
)

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *Todo) Validate() error {
	titlen := utf8.RuneCountInString(t.Title)
	desclen := utf8.RuneCountInString(t.Description)
	if titlen > MaxTitleLength {
		return errors.New("title too long")
	}
	if desclen > MaxDescriptionLength {
		return errors.New("description too long")
	}
	return nil
}
