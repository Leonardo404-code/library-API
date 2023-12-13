package google

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
)

func (c *client) DeleteFromBucket(ctx context.Context, bookTitle string) error {
	titleFormatter := strings.ReplaceAll(bookTitle, " ", "-")
	object := c.storage.Bucket(c.bucketName).Object(titleFormatter)

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("error in get object atributes: %v", err)
	}

	object = object.If(storage.Conditions{
		GenerationMatch: attrs.Generation,
	})

	if err := object.Delete(ctx); err != nil {
		return fmt.Errorf("failed in delete book: %s", bookTitle)
	}

	return nil
}
