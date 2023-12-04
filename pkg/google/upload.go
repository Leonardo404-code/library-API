package google

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	libraryPkg "library-api/pkg/library"
)

func (c *client) Upload(
	ctx context.Context,
	bookInfo *libraryPkg.Book,
	bookFile multipart.File,
) error {
	titleFormatter := strings.ReplaceAll(bookInfo.Title, " ", "-")

	writer := c.storage.Bucket(c.bucketName).Object(titleFormatter).NewWriter(ctx)
	defer writer.Close()

	if writer == nil {
		return fmt.Errorf("failed in create a new writer for book %s", bookInfo.Title)
	}

	if _, err := io.Copy(writer, bookFile); err != nil {
		return fmt.Errorf("error in copy file content to writer: %v", err)
	}

	c.logger.Info("book %s successfully uploaded", bookInfo.Title)

	return nil
}
