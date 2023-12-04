package google

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"library-api/internal/env"
	"library-api/pkg/logger"
)

func New() (Google, error) {
	storageClient, err := storage.NewClient(
		context.Background(),
		option.WithCredentialsFile(credentialPath),
	)
	if err != nil {
		return nil, err
	}

	bucketName := env.GetString(BUCKETNAME)
	bucketURL := env.GetString(BUCKETURL)
	newColorful := logger.NewColorful(logger.White, logger.Yellow, logger.Red)

	return &client{
		storage:    storageClient,
		bucketName: bucketName,
		bucketURL:  bucketURL,
		logger:     newColorful,
	}, nil
}
