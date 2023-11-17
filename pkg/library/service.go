package library

import (
	"mime/multipart"
)

type Services interface {
	CreateBook(*Book, multipart.File) (*Book, error)
	UploadBook(multipart.File) error
}
