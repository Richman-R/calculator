package domain

import "errors"

var (
	ErrUnknownOperation = errors.New("unknown operation")
	ErrDivisionByZero   = errors.New("division by zero")
)
