package google

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"library-api/internal/env"
	"library-api/pkg/logger"
)

func New() (Google, error) {
	credentialsPath := env.GetString(CREDENTIALSPATH)

	storageClient, err := storage.NewClient(
		context.Background(),
		option.WithCredentialsFile(credentialsPath),
	)
	if err != nil {
		return nil, err
	}

	bucketName := env.GetString(BUCKETNAME)
	newColorful := logger.NewColorful(logger.White, logger.Yellow, logger.Red)

	return &client{
		storage:    storageClient,
		bucketName: bucketName,
		logger:     newColorful,
	}, nil
}
