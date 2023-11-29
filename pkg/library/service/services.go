package service

import (
	"library-api/pkg/google"
	"library-api/pkg/library"
)

type service struct {
	libraryRepo library.Repository
	storage     google.Google
}

func Must(
	libraryRepo library.Repository,
	storage google.Google,
) library.Services {
	svc := &service{
		libraryRepo: libraryRepo,
		storage:     storage,
	}

	return svc
}
