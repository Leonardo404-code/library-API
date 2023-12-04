package google

import (
	"cloud.google.com/go/storage"

	"library-api/pkg/logger"
)

type client struct {
	storage    *storage.Client
	logger     logger.Logger
	bucketName string
	bucketURL  string
}
