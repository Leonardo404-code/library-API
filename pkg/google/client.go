package google

import (
	"context"
	"library-api/internal/env"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func New() (Google, error) {
	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, err
	}

	bucketName := env.GetString(BUCKETNAME)
	bucketURL := env.GetString(BUCKETURL)

	return &client{
		storage:    storageClient,
		bucketName: bucketName,
		bucketURL:  bucketURL,
	}, nil
}
