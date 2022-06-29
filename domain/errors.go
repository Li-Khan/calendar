package domain

import "errors"

var (
	ErrAlreadyExist     error = errors.New("event already exist")
	ErrDateAlreadyExist error = errors.New("date already exist")
)
