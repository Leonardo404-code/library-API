package library

import (
	"mime/multipart"
)

type Repository interface {
	Search(filter *Filter) ([]*Book, error)
	CreateBook(bookInfo *Book, bookFile multipart.File, uploadToBucket func() error) error
	DeleteBook(filter *Filter) error
}
