package library

import (
	"mime/multipart"
)

type Repository interface {
	Search(*Filter) ([]*Book, error)
	CreateBook(*Book, multipart.File, func() error) error
}
