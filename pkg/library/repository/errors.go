package repository

import "errors"

var (
	ErrInvalidID   = errors.New("invalid ID")
	ErrFindBook    = errors.New("error in find book")
	ErrParseResult = errors.New("error in parse find result")
)
