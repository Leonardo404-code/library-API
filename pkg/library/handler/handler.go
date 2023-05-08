package handler

import (
	"sync"

	"library-api/pkg/library"
)

type handler struct {
	librarySvc library.Services
}

var (
	h    library.Handlers
	once sync.Once
)

func Must(
	librarySvc library.Services,
) library.Handlers {
	once.Do(func() {
		h = &handler{
			librarySvc: librarySvc,
		}
	})

	return h
}
