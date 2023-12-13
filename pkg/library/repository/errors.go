package repository

import (
	"errors"
)

var (
	ErrFindBook          = errors.New("error in find book")
	ErrParseResult       = errors.New("error in parse find result")
	ErrCreateSession     = errors.New("error in create session")
	ErrStartSession      = errors.New("error in start session")
	ErrStartTransaction  = errors.New("error in start transaction")
	ErrCommitTransaction = errors.New("error in commit transaction")
)
