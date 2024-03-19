package entities

import "errors"

var (
	ErrInvalidInput = errors.New("error -> invalid input(s)")
	ErrLongInput    = errors.New("error -> too long input(s)")
	ErrExists       = errors.New("error -> exists already")
)
