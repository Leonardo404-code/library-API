package library

import (
	"mime/multipart"
)

type Services interface {
	CreateBook(bookInfo *Book, bookFile multipart.File) (*Book, error)
	DeleteBook(requestParams *Filter) error
	DownloadBook(filter *Filter) error
}
