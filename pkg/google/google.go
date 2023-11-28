package google

import (
	"mime/multipart"

	libraryPkg "library-api/pkg/library"
)

type Google interface {
	Upload(bookInfo *libraryPkg.Book, bookFile multipart.File) error
}
