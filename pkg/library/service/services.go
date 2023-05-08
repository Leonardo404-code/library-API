package service

import "library-api/pkg/library"

type service struct {
	libraryRepo library.Repository
}

func Must(
	libraryRepo library.Repository,
) library.Services {
	svc := &service{
		libraryRepo: libraryRepo,
	}

	return svc
}
