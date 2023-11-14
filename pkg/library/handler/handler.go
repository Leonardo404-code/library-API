package handler

import (
	"sync"

	"library-api/pkg/library"
)

type handler struct {
	librarySvc  library.Services
	libraryRepo library.Repository
}

var (
	h    library.Handlers
	once sync.Once
)

func Must(
	librarySvc library.Services,
	libraryRepo library.Repository,
) library.Handlers {
	once.Do(func() {
		h = &handler{
			librarySvc:  librarySvc,
			libraryRepo: libraryRepo,
		}
	})

	return h
}
