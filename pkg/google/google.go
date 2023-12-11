package google

import (
	"context"
	"mime/multipart"

	libraryPkg "library-api/pkg/library"
)

type Google interface {
	Upload(ctx context.Context, bookInfo *libraryPkg.Book, bookFile multipart.File) error
	DownloadBook(ctx context.Context, objectTitle string) error
	DeleteFromBucket(ctx context.Context, objectTitle string) error
}
