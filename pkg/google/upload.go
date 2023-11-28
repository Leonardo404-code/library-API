package google

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
	"time"

	libraryPkg "library-api/pkg/library"
)

func (c *client) Upload(bookInfo *libraryPkg.Book, bookFile multipart.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	titleFormatter := strings.ReplaceAll(bookInfo.Title, " ", "-")

	writer := c.storage.Bucket(c.bucketName).Object(titleFormatter).NewWriter(ctx)
	defer writer.Close()

	if writer == nil {
		return fmt.Errorf("failed to get writer for object %s", bookInfo.Title)
	}

	if _, err := io.Copy(writer, bookFile); err != nil {
		return fmt.Errorf("error in copy file to writer: %v", err)
	}

	bookInfo.BookURL = fmt.Sprintf("%s/%s", c.bucketURL, titleFormatter)

	return nil
}
