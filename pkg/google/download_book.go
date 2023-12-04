package google

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

func (c *client) DownloadBook(ctx context.Context, bookTitle string) error {
	filename := fmt.Sprintf("%s.pdf", bookTitle)

	bookFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error in create file: %v", err)
	}

	titleFormatter := strings.ReplaceAll(bookTitle, " ", "-")

	bookContent, err := c.storage.Bucket(c.bucketName).Object(titleFormatter).NewReader(ctx)
	if err != nil {
		return fmt.Errorf(`error in read %s object: %v`, bookTitle, err)
	}

	defer bookContent.Close()

	if _, err := io.Copy(bookFile, bookContent); err != nil {
		return fmt.Errorf("error in copy object content to file: %v", err)
	}

	if err = bookFile.Close(); err != nil {
		return fmt.Errorf("error in close file: %v", err)
	}

	c.logger.Info("download of book %s is complete", bookTitle)

	return nil
}
