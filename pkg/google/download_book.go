package google

import (
	"context"
	"fmt"
	"io"
	"os"
)

func (c *client) DownloadBook(ctx context.Context, bookTitle string) error {
	filename := fmt.Sprintf("%s.pdf", bookTitle)
	bookFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error in create file: %v", err)
	}

	rc, err := c.storage.Bucket(c.bucketName).Object(bookTitle).NewReader(ctx)
	if err != nil {
		return fmt.Errorf(`error in create a new write for object %s: %v`, bookTitle, err)
	}
	defer rc.Close()

	if _, err := io.Copy(bookFile, rc); err != nil {
		return fmt.Errorf("error in copy RC to file: %v", err)
	}

	if err = bookFile.Close(); err != nil {
		return fmt.Errorf("error in close file: %v", err)
	}

	fmt.Printf("Blob %s downloaded to local\n", bookTitle)

	return nil
}
