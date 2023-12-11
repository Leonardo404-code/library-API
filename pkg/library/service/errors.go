package service

import (
	"errors"
)

var (
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidDate        = errors.New("invalid date")
	ErrCreateBook         = errors.New("error in create book")
	ErrInvalidObjectID    = errors.New("invalid objectID")
	ErrGetBook            = errors.New("error in get book")
	ErrNotFound           = errors.New("error books not found")
	ErrDeleteBook         = errors.New("error in delete book")
)
