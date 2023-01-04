package service

import "library-api/pkg/library"

type services struct {
	libraryRepo library.Repository
}

func Must(
	libraryRepo library.Repository,
) library.Services {
	svc := &services{
		libraryRepo: libraryRepo,
	}

	return svc
}
